package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	// "fmt"
)

// 项目控制器
type Target struct{ App }

func (c Target) PostCreate() r.Result {
	return c.Render()
}

func (c Target) Show() r.Result {
	return c.Render()
}

func (c Target) Create() r.Result {
	return c.Render()
}
