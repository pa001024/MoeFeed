package controllers

import (
	"net/url"
	"regexp"

	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	"github.com/pa001024/MoeFeed/app/service"
	"github.com/pa001024/MoeWorker/util"
	r "github.com/robfig/revel"
)

// 用户账户控制器
type Account struct{ CommonDomain }

/*
SSO单点登录 原理与流程

客户方: 需要登录的网站/客户端
服务方: 本站

1. 客户方通过跳转(302)提交token给服务方
2. 服务方通过token判断客户身份并记录在session中
3. 如果该token不存在 服务方跳转到登录界面 用户输入密码进行验证
4. 如果验证通过或者token已存在则进入下一步 否则让用户再次输入密码
5. 服务端通过跳转将accesscode交给客户端
6. 客户方使用secretcode获取data后进行内部登录

*/

// [动][边] SSO登录
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
	user := po.GetAccountByNameOrEmail(username)
	if user == nil {
		c.Flash.Error(c.Message("user.notexists"))
		return c.Redirect("/login?return_to=%s", return_to)
	}
	r.INFO.Printf("[DoLogin] %+v", user)
	// 验证密码
	if !user.ValidatePassword(password) {
		c.Validation.Error(c.Message("password.error"))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/login?return_to=%s", return_to)
	}
	// 登陆成功写入session
	c.Session[USER] = util.ToString(user.Id)
	// 单点登录
	token := po.GenerateToken(user.Id)
	c.Session[TOKEN] = token.Token

	if return_to == "" {
		return_to = "/"
	}

	return c.Redirect(return_to)
}

var (
	_accountRegex = regexp.MustCompile(`^[A-z0-9_\-]*$`)
	_emailRegex   = regexp.MustCompile(`\w+(?:[-+.]\w+)*@\w+(?:[-.]\w+)*\.\w+(?:[-.]\w+)*`)
)

// [动][写] 用户注册
func (c *Account) DoRegister(user *models.Account, return_to, password string) r.Result {
	if c.Validation.Required(user.Email).Message(c.Message("email.required")).Ok {
		c.Validation.Check(user.Email, r.MinSize{6}, r.MaxSize{100}, r.Match{_emailRegex}).
			Message(c.Message("email.invalid"))
	}
	if c.Validation.Required(user.Username).Message(c.Message("username.required")).Ok {
		c.Validation.Check(user.Username, r.MinSize{2}, r.MaxSize{32}, r.Match{_accountRegex}).
			Message(c.Message("username.invalid"))
	}
	if c.Validation.Required(password).Message(c.Message("password.required")).Ok {
		c.Validation.Check(password, r.MaxSize{32}, r.MinSize{6}).
			Message(c.Message("password.invalid"))
	}

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/register")
	}
	po := repo.CommonRepo()
	defer po.Close()
	// 检查用户是否已存在
	if po.GetAccountByName(user.Username) != nil {
		c.Flash.Error(c.Message("user.exists"))
		return c.Redirect("/register?return_to=%s", return_to)
	}
	// 创建用户
	user.Password = user.GeneratePassword(password)
	user.Enable = false // 安全
	po.Put(user)
	r.INFO.Println("[NewAccount]", user)
	// 写入code
	code := po.GenerateEmailVerfiy(user.Id)
	// 发送邮件
	r.INFO.Println("[NewAccountEmailVerfiy]", user.Email)
	aurl := "http://localhost:9000/reauth?" + (url.Values{"id": {user.Username}, "code": {code.Code}}).Encode()
	service.Mail.SendMail(user.Email, "完成你的注册", `
	<p>如果你需要使用全部功能 请点击下列链接完成验证</p>
	<p><a href="`+aurl+`">`+aurl+`</a></p>`) // TODO: i18n

	c.Session[USER] = util.ToString(user.Id)
	if return_to == "" {
		return_to = "/"
	}
	return c.Redirect(return_to)
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

// [静] 用户登入前端
func (c *Account) Login(return_to string) r.Result {
	// 已登录则跳转到首页 TODO: 需要改进
	if c.CheckAccountAndClose() != nil {
		return c.Redirect("/")
	}
	return c.Render(return_to)
}

// [静] 用户登出前端
func (c *Account) Logout() r.Result {
	delete(c.Session, USER)
	return c.Redirect("/")
}

// [静] 用户注册前端
func (c *Account) Register(return_to string) r.Result {
	if return_to == "" && c.Request.Referer() != "" {
		return_to = c.Request.Referer()
	}
	return c.Render(return_to)
}
