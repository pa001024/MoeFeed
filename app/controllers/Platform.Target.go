package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	r "github.com/robfig/revel"
)

// 目标点控制器
type Target struct{ PlatformDomain }

// [动][写]
func (c Target) DoCreate(user, project string, target *models.Target) r.Result {
	u, p, po := c.CheckAccessProjectRenderArgs(user, project)
	if u == nil {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/p/%s/%s", user, project)
	}
	if p == nil {
		return c.NotFound(c.Message("project.notfound"))
	}
	target.ProjectId = p.Id
	po.Put(target)
	return c.Redirect("/p/%s/%s", user, project)
}

// [静]显示单个目标点
func (c Target) Show(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mReadable"].(bool) {
		c.Flash.Error(c.Message("project.view.nopermission"))
		return c.Redirect("/p/%s/%s", user, project)
	}
	return c.Render()
}

// [静]创建前端
func (c Target) Create(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mEditable"].(bool) {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/p/%s/%s", user, project)
	}
	return c.Render()
}
