package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	// "fmt"
)

// 项目控制器
type Source struct{ App }

func (c Source) PostCreate() r.Result {
	return c.Render()
}

func (c Source) Show() r.Result {
	return c.Render()
}

func (c Source) Create() r.Result {
	return c.Render()
}
