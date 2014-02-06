package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
)

// 后台管理
type AdminDomain struct{ CommonDomain }

// 管理员身份验证持久化
func (c *AdminDomain) CheckAdmin() (u *models.Account, po *repo.Common) {
	if accountId, ok := c.Session[USER]; ok {
		po = repo.CommonRepo()
		u = po.GetAccount(accountId)
		// if u.Type == models.AccountAdmin {
		// 	c.RenderArgs["mAdmin"] = u
		// 	return
		// }
	}
	return
}

// 管理员身份验证持久化
func (c *AdminDomain) CheckAdminAndClose() (u *models.Account) {
	u, po := c.CheckAdmin()
	if po != nil {
		po.Close()
	}
	return u
}
