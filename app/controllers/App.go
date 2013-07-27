package controllers

import (
	"fmt"
	"log"
	"net/url"

	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	"github.com/pa001024/MoeFeed/app/service"
	"github.com/pa001024/MoeWorker/util"
	r "github.com/robfig/revel"
)

const (
	USER = "user"
)

// 基础应用
type App struct{ *r.Controller }

// 用户状态持久化
func (c App) CheckUser() *models.User {
	if vu := c.RenderArgs["mUser"]; vu != nil {
		return vu.(*models.User)
	}
	if userId, ok := c.Session[USER]; ok {
		u := repo.UserRepo.GetById(userId)
		c.RenderArgs["mUser"] = u
		return u
	}
	return nil
}

// [动][边] 登录
func (c App) PostLogin(username, password, return_to string) r.Result {
	c.Validation.Required(username).Message("请填写用户名")
	c.Validation.Required(password).Message("请填写密码")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/login?return_to=%s", return_to)
	}
	user := repo.UserRepo.GetByNameOrEmail(username)
	if user == nil {
		c.Flash.Error("用户不存在")
		return c.Redirect("/login?return_to=%s", return_to)
	}
	// 验证密码
	user.ValidatePassword(c.Validation, password)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/login?return_to=%s", return_to)
	}
	// 登陆成功写入session
	c.Session[USER] = util.ToString(user.Id)

	if return_to == "" {
		return_to = "/"
	}

	return c.Redirect(return_to)
}

// [动][写] 用户注册
func (c App) PostRegister(user *models.User, return_to, password, password2 string) r.Result {
	c.Validation.Required(password == password2).
		Message("两次密码不匹配")
	user.Validate(c.Validation, password)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/register")
	}
	// 创建用户
	user.Password = user.GeneratePassword(password)
	fmt.Printf("[NewUser] %#v\n", user)
	repo.UserRepo.Put(user)
	// 写入code
	code := &models.UserCode{UserId: user.Id}
	code.GenerateCode()
	repo.UserCodeRepo.Put(code)
	// 发送邮件
	log.Println("Send mail to", user.Email)
	aurl := "http://feed.qaq.ca/reauth?" + (url.Values{"id": {user.Username}, "code": {code.Code}}).Encode()
	service.Mail.SendMailAsync(user.Email, "完成你的注册", `
		<p>如果你需要使用全部功能 请点击下列链接完成验证</p>
		<p><a href="`+aurl+`">`+aurl+`</a></p>
		`)

	c.Session[USER] = fmt.Sprint(user.Id)
	if return_to == "" {
		return_to = "/"
	}
	return c.Redirect(return_to)
}

// [动][写] 验证
func (c App) Reauth(code string, id int64) r.Result {
	msg := "验证失败"
	usercode := repo.UserCodeRepo.GetByOwnerAndCode(code, id)
	if usercode != nil && usercode.User != nil {
		if usercode.User.Status == models.UnauthedUser {
			usercode.User.Status = models.AuthedUser
			repo.UserRepo.Put(usercode.User)
			repo.UserCodeRepo.Delete(usercode)
			msg = "您的账户已通过验证"
		} else if usercode.User.Status == models.UnauthedTeam {
			usercode.User.Status = models.AuthedTeam
			repo.UserRepo.Put(usercode.User)
			repo.UserCodeRepo.Delete(usercode)
			msg = "您的组织账户已通过验证"
		} else {
			msg = "您无需进行验证"
		}
	}
	return c.Render(msg)
}

// [静] 主页
func (c App) Index() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静] 帮助
func (c App) Help() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静] 用户登入前端
func (c App) Login(return_to string) r.Result {
	// 已登录则跳转到首页
	if c.CheckUser() != nil {
		return c.Redirect("/")
	}
	return c.Render(return_to)
}

// [静] 用户登出前端
func (c App) Logout() r.Result {
	delete(c.Session, USER)
	return c.Redirect("/")
}

// [静] 用户注册前端
func (c App) Register(return_to string) r.Result {
	if return_to == "" && c.Request.Referer() != "" {
		return_to = c.Request.Referer()
	}
	return c.Render(return_to)
}

func assetsError(err error) {
	if err != nil {
		panic(err)
	}
}
