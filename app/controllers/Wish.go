package controllers

import (
	// "github.com/pa001024/MoeFeed/app/models"
	// repo "github.com/pa001024/MoeFeed/app/repository"
	// "github.com/pa001024/MoeFeed/app/service"
	// "github.com/pa001024/MoeWorker/util"
	r "github.com/robfig/revel"
)

// 用户对功能的愿景
type Wish struct{ App }

func (c Wish) Show() r.Result {
	return c.Todo()
}
