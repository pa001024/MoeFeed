package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	// "fmt"
)

// 项目控制器
type Callback struct{ App }

func (c Callback) PostCreate() r.Result {
	return c.Render()
}

func (c Callback) Show() r.Result {
	return c.Render()
}

func (c Callback) Create() r.Result {
	return c.Render()
}
