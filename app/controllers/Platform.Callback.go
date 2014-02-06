package controllers

import (
	"fmt"

	"github.com/pa001024/MoeFeed/app/models"
	r "github.com/robfig/revel"
)

// 回调控制器
type Callback struct{ PlatformDomain }

func (c Callback) DoCreate(user, project string, callback *models.Callback) r.Result {
	u, p, po := c.CheckAccessProjectRenderArgs(user, project)
	defer po.Close()
	if u == nil {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	if u == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/%s/%s", user, project)
	}
	callback.ProjectId = p.Id
	po.Put(callback)
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
	p, po := c.CheckProject(user, project)
	defer po.Close()
	b := po.GetCallback(url, p.Id)
	trueurl := fmt.Sprintf("/%s/%s/callback/%s", user, project, url)
	if b == nil {
		return c.RenderJson(CallbackError{trueurl, 404, "找不到此回调"})
	}
	// TODO: 添加实际调用
	return c.RenderJson(CallbackError{trueurl, 200, "成功"})
}

func (c Callback) Show(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mReadable"].(bool) {
		c.Flash.Error(c.Message("project.view.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	return c.Render()
}

// [静]创建前端
func (c Callback) Create(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mEditable"].(bool) {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	return c.Render()
}
