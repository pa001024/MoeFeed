package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 资源控制器
type Resource struct{ Project }

// [动][写]
func (c Resource) PostCreate(user, project string, source *models.Resource) r.Result {
	u := c.CheckUser()
	p := c.CheckProject(user, project)
	if p == nil {
		return c.NotFound("找不到该项目")
	}
	if u == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/%s/%s", user, project)
	}
	repo.ResourceRepo.Put(source)
	return c.Redirect("/%s/%s", user, project)
}

// [静]显示单个资源
func (c Resource) Show(user, project string) r.Result {
	return c.Render()
}

// [静]创建前端
func (c Resource) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
