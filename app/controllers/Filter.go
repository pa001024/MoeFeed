package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	// "fmt"
)

// 项目控制器
type Filter struct{ App }

func (c Filter) PostCreate() r.Result {
	return c.Render()
}

func (c Filter) Show() r.Result {
	return c.Render()
}

func (c Filter) Create() r.Result {
	return c.Render()
}
