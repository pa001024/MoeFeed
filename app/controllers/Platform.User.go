package controllers

import (
	r "github.com/robfig/revel"
)

// 用户控制器
type PlatformUser struct{ PlatformDomain }

// 跳转
func (c *PlatformUser) ProfileLink(user string) r.Result {
	return c.Redirect("/%s", user)
}

// 用户基本信息
func (c *PlatformUser) Profile(user string) r.Result {
	// c.CheckUser()
	return c.Todo()
}

// 用户安全信息
func (c *PlatformUser) Security(user string) r.Result {
	// c.CheckUser()
	return c.Todo()
}

// [静]用户展示页
func (c *PlatformUser) Show(user string) r.Result {
	_, po := c.CheckUser()
	defer po.Close()
	mUser := po.GetUser(user)
	if mUser == nil {
		return c.NotFound(c.Message("user.notfound"))
	}
	mProjects := po.FindProjectByOwner(mUser.Id)
	return c.Render(mUser, mProjects)
}

// [静]平台用户登录前端
func (c *PlatformUser) Login(return_to string) r.Result {
	return c.Render()
}
