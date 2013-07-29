package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 帮助页面
type Help struct{ App }

// [静] 帮助
func (c Help) Index() r.Result {
	c.CheckUser()
	return c.Render()
}
