package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 目标点控制器
type Target struct{ Project }

// [动][写]
func (c Target) DoCreate(user, project string, target *models.Target) r.Result {
	u, p := c.CheckOwnerProject(user, project)
	if u == nil {
		c.Flash.Error("你没有权限编辑该项目")
		return c.Redirect("/%s/%s", user, project)
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
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}

// [静]创建前端
func (c Target) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
