package controllers

import (
	"net/url"

	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	"github.com/pa001024/MoeFeed/app/service"
	"github.com/pa001024/MoeWorker/util"
	r "github.com/robfig/revel"
)

// 用户控制器
type User struct{ App }

// 跳转
func (c User) ProfileLink(user string) r.Result {
	return c.Redirect("/%s", user)
}

// 用户基本信息
func (c User) Profile(user string) r.Result {
	// c.CheckUser()
	return c.Todo()
}

// 用户安全信息
func (c User) Security(user string) r.Result {
	// c.CheckUser()
	return c.Todo()
}

// [静]用户展示页
func (c User) Show(user string) r.Result {
	c.CheckUser()
	mUser := repo.UserRepo.GetByName(user)
	if mUser == nil {
		return c.NotFound(c.Message("user.notfound"))
	}
	mProjects := repo.ProjectRepo.FindByOwner(mUser.Id)
	return c.Render(mUser, mProjects)
}

// [动][边] 登录
func (c User) DoLogin(username, password, return_to string) r.Result {
	c.Validation.Required(username).Message(c.Message("username.required"))
	c.Validation.Required(password).Message(c.Message("password.required"))

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/login?return_to=%s", return_to)
	}
	user := repo.UserRepo.GetByNameOrEmail(username)
	if user == nil {
		c.Flash.Error(c.Message("user.notexists"))
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
func (c User) DoRegister(user *models.User, return_to, password string) r.Result {
	user.Validate(c.Validation, password)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/register")
	}
	// 创建用户
	user.Password = user.GeneratePassword(password)
	r.INFO.Println("[NewUser]", user)
	repo.UserRepo.Put(user)
	// 写入code
	code := &models.UserCode{UserId: user.Id}
	code.GenerateCode()
	repo.UserCodeRepo.Put(code)
	// 发送邮件
	r.INFO.Println("Send mail to", user.Email)
	aurl := "http://feed.qaq.ca/reauth?" + (url.Values{"id": {user.Username}, "code": {code.Code}}).Encode()
	service.Mail.SendMail(user.Email, "完成你的注册", `
		<p>如果你需要使用全部功能 请点击下列链接完成验证</p>
		<p><a href="`+aurl+`">`+aurl+`</a></p>
		`)

	c.Session[USER] = util.ToString(user.Id)
	if return_to == "" {
		return_to = "/"
	}
	return c.Redirect(return_to)
}

// [动][写] 验证
func (c User) Reauth(code string, id int64) r.Result {
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

// [静] 用户登入前端
func (c User) Login(return_to string) r.Result {
	// 已登录则跳转到首页
	if c.CheckUser() != nil {
		return c.Redirect("/")
	}
	return c.Render(return_to)
}

// [静] 用户登出前端
func (c User) Logout() r.Result {
	delete(c.Session, USER)
	return c.Redirect("/")
}

// [静] 用户注册前端
func (c User) Register(return_to string) r.Result {
	if return_to == "" && c.Request.Referer() != "" {
		return_to = c.Request.Referer()
	}
	return c.Render(return_to)
}
