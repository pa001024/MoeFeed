package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 回调控制器
type Callback struct{ Project }

func (c Callback) PostCreate(user, project string, callback *models.Callback) r.Result {
	u := c.CheckUser()
	p := c.CheckProject(user, project)
	if p == nil {
		return c.NotFound("找不到该项目")
	}
	if u == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/%s/%s", user, project)
	}
	callback.ProjectId = p.Id
	repo.CallbackRepo.Put(callback)
	return c.Redirect("/%s/%s", user, project)
}

func (c Callback) Show(user, project, callback string) r.Result {
	// c.CheckUser()
	p := c.CheckProject(user, project)
	repo.CallbackRepo.GetByProjectAndUrl(callback, p.Id)
	return c.Render()
}

func (c Callback) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
