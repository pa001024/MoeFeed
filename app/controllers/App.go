package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

const (
	USER = "user"
)

// 基础应用
type App struct{ *r.Controller }

// 用户状态持久化
func (c *App) CheckUser() *models.User {
	if vu := c.RenderArgs["mUser"]; vu != nil {
		return vu.(*models.User)
	}
	if userId, ok := c.Session[USER]; ok {
		u := repo.UserRepo.GetById(userId)
		c.RenderArgs["mUser"] = u
		return u
	}
	return nil
}

func assetsError(err error) {
	if err != nil {
		panic(err)
	}
}
