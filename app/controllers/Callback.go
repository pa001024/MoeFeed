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

// 错误信息
type CallbackError struct {
	Url  string `json:"url"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (c Callback) Call(user, project, url string) r.Result {
	// c.CheckUser()
	p := c.CheckProject(user, project)
	b := repo.CallbackRepo.GetByProjectAndUrl(url, p.Id)
	if b == nil {
		return c.RenderJson(CallbackError{"/" + user + "/" + project + "/callback/" + url, 404, "找不到此回调"})
	}
	// TODO: 添加实际调用
	return c.RenderJson(CallbackError{"/" + user + "/" + project + "/callback/" + url, 200, "成功"})
}

func (c Callback) Show(user, project string) r.Result {
	// c.CheckUser()
	// c.CheckProject(user, project)
	return c.Redirect("/%s/%s", user, project)
}

func (c Callback) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
