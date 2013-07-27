package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	// "fmt"
)

// 项目控制器
type Resource struct{ App }

func (c Resource) PostCreate() r.Result {
	return c.Render()
}

func (c Resource) Show() r.Result {
	return c.Render()
}

func (c Resource) Create() r.Result {
	return c.Render()
}
