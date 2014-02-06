package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	r "github.com/robfig/revel"
)

// 过滤器控制器
type Filter struct{ PlatformDomain }

// [动][写]
func (c Filter) DoCreate(user, project string, filter *models.Filter) r.Result {
	u, p, po := c.CheckAccessProjectRenderArgs(user, project)
	defer po.Close()
	if p == nil {
		return c.NotFound(c.Message("project.notfound"))
	}
	if u == nil {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	filter.ProjectId = p.Id
	po.Put(filter)
	return c.Redirect("/%s/%s", user, project)
}

// [静]显示单个过滤器
func (c Filter) Show(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mReadable"].(bool) {
		c.Flash.Error(c.Message("project.view.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	return c.Render()
}

// [静]创建前端
func (c Filter) Create(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mEditable"].(bool) {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	return c.Render()
}
