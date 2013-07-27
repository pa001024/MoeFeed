package controllers

import (
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

type User struct {
	App
}

// 跳转
func (c User) ProfileLink() r.Result {
	return c.Redirect("/")
}

// 用户基本信息
func (c User) Profile() r.Result {
	c.CheckUser()
	return c.Render()
}

// 用户安全信息
func (c User) Security() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静]用户展示页
func (c User) Show(user string) r.Result {
	c.CheckUser()
	mUser := repo.UserRepo.GetByName(user)
	if mUser == nil {
		return c.NotFound("没有此用户")
	}
	mProjects := repo.ProjectRepo.FindByOwner(mUser.Id)
	return c.Render(mUser, mProjects)
}
