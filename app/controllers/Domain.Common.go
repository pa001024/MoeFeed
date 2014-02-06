package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

type CommonDomain struct{ *r.Controller }

func (c *CommonDomain) CheckAccount() (u *models.Account, po *repo.Common) {
	if id, ok := c.Session[ACCOUNT]; ok {
		po = repo.CommonRepo()
		u = po.GetAccount(id)
		c.RenderArgs["mAccount"] = u
	}
	return
}
func (c *CommonDomain) CheckAccountAndClose() *models.Account {
	u, po := c.CheckAccount()
	if po != nil {
		po.Close()
	}
	return u
}
