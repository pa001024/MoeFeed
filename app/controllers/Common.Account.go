package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	"github.com/pa001024/MoeFeed/app/service"
	r "github.com/robfig/revel"
)

// 总账户接口
// 姿势:
// BrowserID(persona) -> Account -> AppUser
type Account struct{ CommonDomain }

const _VERIFY_URL = "https://verifier.login.persona.org/verify"

var (
	_accountRegex = regexp.MustCompile(`^[A-z0-9_\-]*$`)
	_emailRegex   = regexp.MustCompile(`\w+(?:[-+.]\w+)*@\w+(?:[-.]\w+)*\.\w+(?:[-.]\w+)*`)
)

// BrowserID协议登录
func (c *Account) DoLoginBrowserID(assertion, return_to string) r.Result {
	res, err := http.Post(_VERIFY_URL, "application/x-www-form-urlencoded;charset=UTF-8",
		strings.NewReader(fmt.Sprint(`assertion=`, assertion, `&audience=http://localhost:9000`)), // TODO: 需配置
	)
	if err != nil {
		return c.RenderError(err)
	}
	var verifyResult struct {
		Status   string `json:"status"`
		Email    string `json:"email"`
		Audience string `json:"audience"`
		Expires  uint64 `json:"expires"`
		Issuer   string `json:"issuer"`
	}
	err = json.NewDecoder(res.Body).Decode(&verifyResult)
	res.Body.Close()
	if err != nil {
		return c.RenderError(err)
	}
	r.INFO.Println(verifyResult)
	if verifyResult.Status == "okay" {
		po := repo.CommonRepo()
		account := po.GetAccountByEmail(verifyResult.Email)
		po.Close()
		// 新用户跳转到signup页面
		if account == nil {
			c.Session[EMAIL] = verifyResult.Email
			return c.Redirect("/signup?return_to=%s", return_to)
		} else {
			c.SetLogin(account)
		}

		// 登录成功跳转到return_to
		if return_to == "" {
			return_to = "/"
		}
		return c.Redirect("%s", return_to)
	}
	c.Response.Status = 400
	return c.RenderText(c.Message("login.failed"))
}

// [动] 用户登入
func (c *Account) DoLogin(username, password, return_to string) r.Result {
	// 必填字段
	c.Validation.Required(username).Message(c.Message("username.required"))
	c.Validation.Required(password).Message(c.Message("password.required"))

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/login?return_to=%s", return_to)
	}
	po := repo.CommonRepo()
	defer po.Close()
	account := po.GetAccountByNameOrEmail(username)
	if account == nil {
		c.Flash.Error(c.Message("user.notexists"))
		return c.Redirect("/login?return_to=%s", return_to)
	}
	r.INFO.Printf("[DoLogin] %+v", account)
	// 验证密码
	if !account.ValidatePassword(password) {
		c.Validation.Error("password.error")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/login?return_to=%s", return_to)
	}
	// 登陆成功写入session
	c.SetLogin(account)

	if return_to == "" {
		return_to = "/"
	}
	return c.Redirect(return_to)
}

// [静] 用户登入前端
func (c *Account) Login(return_to string) r.Result {
	// 已登录则跳转到首页 TODO: 需要改进
	if c.CheckAccountAndClose() != nil {
		return c.Redirect("/")
	}
	return c.Render(return_to)
}

// [动][写] 注册
func (c *Account) DoSignup(account *models.Account, password, return_to string) r.Result {
	// 防改
	if c.Session[EMAIL] != "" {
		account.Enable = true
		account.Email = c.Session[EMAIL]
	} else {
		account.Enable = false
	}
	// 防乱
	if c.Validation.Required(account.Email).Message(c.Message("email.required")).Ok {
		c.Validation.Check(account.Email, r.MinSize{6}, r.MaxSize{100}, r.Match{_emailRegex}).
			Message(c.Message("email.invalid"))
	}
	if c.Validation.Required(account.Username).Message(c.Message("username.required")).Ok {
		c.Validation.Check(account.Username, r.MinSize{2}, r.MaxSize{32}, r.Match{_accountRegex}).
			Message(c.Message("username.invalid"))
	}
	if c.Validation.Required(password).Message(c.Message("password.required")).Ok {
		c.Validation.Check(password, r.MaxSize{32}, r.MinSize{6}).
			Message(c.Message("password.invalid"))
	}
	po := repo.CommonRepo()
	// 防撞
	if po.GetAccountByName(account.Username) != nil {
		po.Close()
		return c.RenderError(fmt.Errorf("%s", c.Message("user.exists")))
	}
	// 密码可选
	if password != "" {
		account.GeneratePassword(password)
	}
	po.Put(account)
	// 防刷
	if account.Enable == false {
		r.INFO.Println("[NewAccount]", account)
		// 写入code
		code := po.GenerateEmailVerfiy(account.Id)
		// 发送邮件
		r.INFO.Println("[NewAccountEmailVerfiy]", account.Email)
		aurl := "http://localhost:9000/reauth?" + (url.Values{"id": {account.Username}, "code": {code.Code}}).Encode()
		service.Mail.SendMail(account.Email, "完成你在MoeFeed的注册", `<p>请点击下列链接完成验证</p><p><a href="`+aurl+`">`+aurl+`</a></p><p>如非本人操作请忽略本邮件</p>`) // TODO: i18n
	}
	po.Close()
	return c.Redirect("%s", return_to)
}

// [静] 注册前端
func (c *Account) Signup(return_to string) r.Result {
	mEmail := c.Session[EMAIL]
	return c.Render(mEmail, return_to)
}

// [动][写] 验证
func (c *Account) Reauth(code string, id string) r.Result {
	msg := c.Message("emailverfiy.fail")
	po := repo.CommonRepo()
	defer po.Close()
	usercode := po.GetEmailVerfiy(code, id)
	if usercode != nil && usercode.Account != nil {
		r.INFO.Println("[Reauth]", usercode)
		if !usercode.Account.Enable {
			usercode.Account.Enable = true
			po.Put(usercode.Account)
			po.Delete(usercode)
			msg = c.Message("emailverfiy.success")
		} else {
			po.Delete(usercode)
			msg = c.Message("emailverfiy.notneed")
		}
	}
	return c.Render(msg)
}

// [动][边] 用户登出
func (c *Account) Logout(return_to string) r.Result {
	c.SetLogout()
	if return_to == "" {
		return_to = "/"
	}
	return c.Redirect("%s", return_to)
}
