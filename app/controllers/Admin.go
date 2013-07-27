package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 后台管理
type Admin struct{ App }

// 用户状态持久化+管理员身份验证
func (c Admin) CheckUser() *models.User {
	if vu := c.RenderArgs["mUser"]; vu != nil {
		return vu.(*models.User)
	}
	if userId, ok := c.Session[USER]; ok {
		u := repo.UserRepo.GetById(userId)
		if u.Status == models.SysAdmin {
			c.RenderArgs["mUser"] = u
			return u
		}
	}
	return nil
}

func (c Admin) Index() r.Result {
	if c.CheckUser() == nil {
		return c.Redirect("/")
	}
	return c.Render()
}
