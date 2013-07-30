package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 过滤器控制器
type Filter struct{ Project }

// [动][写]
func (c Filter) DoCreate(user, project string, filter *models.Filter) r.Result {
	u, p := c.CheckEditableProject(user, project)
	if u == nil {
		c.Flash.Error("你没有权限编辑该项目")
		return c.Redirect("/%s/%s", user, project)
	}
	if u == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/%s/%s", user, project)
	}
	filter.ProjectId = p.Id
	repo.FilterRepo.Put(filter)
	return c.Redirect("/%s/%s", user, project)
}

// [静]显示单个过滤器
func (c Filter) Show(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}

// [静]创建前端
func (c Filter) Create(user, project string) r.Result {
	u, _ := c.CheckEditableProject(user, project)
	if u == nil {
		c.Flash.Error("你没有权限编辑该项目")
		return c.Redirect("/%s/%s", user, project)
	}
	return c.Render()
}
