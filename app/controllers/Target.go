package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 目标点控制器
type Target struct{ Project }

// [动][写]
func (c Target) PostCreate(user, project string, target *models.Target) r.Result {
	u := c.CheckUser()
	p := c.CheckProject(user, project)
	if p == nil {
		return c.NotFound("找不到该项目")
	}
	if u == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/%s/%s", user, project)
	}
	target.ProjectId = p.Id
	repo.TargetRepo.Put(target)
	return c.Redirect("/%s/%s", user, project)
}

// [静]显示单个目标点
func (c Target) Show(user, project string) r.Result {
	return c.Render()
}

// [静]创建前端
func (c Target) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
