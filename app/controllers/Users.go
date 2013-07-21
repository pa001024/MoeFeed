package controllers

import (
	repo "github.com/pa001024/MoeFeed/app/repository"
	"github.com/pa001024/MoeFeed/app/routes"
	r "github.com/robfig/revel"
)

type Users struct {
	App
}

// 跳转
func (c Users) ProfileLink() r.Result {
	return c.Redirect(routes.App.Index())
}

// 用户基本信息
func (c Users) Profile() r.Result {
	c.CheckUser()
	return c.Render()
}

// 用户安全信息
func (c Users) Security() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静]用户展示页
func (c Users) Show(user string) r.Result {
	c.CheckUser()
	mUser := repo.UserRepo.GetByName(user)
	if mUser == nil {
		return c.Redirect("/")
	}
	mProjects := repo.ProjectRepo.FindByOwner(mUser.Id)
	return c.Render(mUser, mProjects)
}
