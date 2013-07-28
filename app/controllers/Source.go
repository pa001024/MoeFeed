package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 来源点控制器
type Source struct{ Project }

// [动][写]
func (c Source) PostCreate(user, project string, source *models.Source) r.Result {
	u := c.CheckUser()
	p := c.CheckProject(user, project)
	if p == nil {
		return c.NotFound("找不到该项目")
	}
	if u == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/%s/%s", user, project)
	}
	source.ProjectId = p.Id
	repo.SourceRepo.Put(source)
	return c.Redirect("/%s/%s", user, project)
}

// [静]显示单个来源点
func (c Source) Show(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}

// [静]创建前端
func (c Source) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
