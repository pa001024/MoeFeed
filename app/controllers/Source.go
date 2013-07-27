package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	// "fmt"
)

// 项目控制器
type Source struct{ Project }

// [动][写]
func (c Source) PostCreate() r.Result {
	return c.Render()
}

// [静]显示单个来源点
func (c Source) Show(user, project string) r.Result {
	return c.Render()
}

// [静]创建前端
func (c Source) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
