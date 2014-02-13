package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	r "github.com/robfig/revel"
)

// 来源点控制器
type Source struct{ PlatformDomain }

// [动][写]
func (c Source) DoCreate(user, project string, source *models.Source) r.Result {
	u, p, po := c.CheckAccessProjectRenderArgs(user, project)
	defer po.Close()
	if u == nil {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/p/%s/%s", user, project)
	}
	if p == nil {
		return c.NotFound(c.Message("project.notfound"))
	}
	source.ProjectId = p.Id
	po.Put(source)
	return c.Redirect("/p/%s/%s", user, project)
}

// [静]显示单个来源点
func (c Source) Show(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mReadable"].(bool) {
		c.Flash.Error(c.Message("project.view.nopermission"))
		return c.Redirect("/p/%s/%s", user, project)
	}
	return c.Render()
}

// [静]创建前端
func (c Source) Create(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mEditable"].(bool) {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/p/%s/%s", user, project)
	}
	return c.Render()
}
