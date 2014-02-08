package controllers

import (
	"fmt"

	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 边缘层
type CommonDomain struct{ *r.Controller }

// 账号状态
func (c *CommonDomain) CheckAccount() (u *models.Account, po *repo.Common) {
	if id, ok := c.Session[ACCOUNTID]; ok {
		po = repo.CommonRepo()
		u = po.GetAccount(id)
		c.RenderArgs["mAccount"] = u
	}
	return
}

// 账号状态
func (c *CommonDomain) CheckAccountAndClose() *models.Account {
	u, po := c.CheckAccount()
	if po != nil {
		po.Close()
	}
	return u
}

// 返回Referer页或者return_to页
func (c *CommonDomain) Return(return_to string) r.Result {
	if return_to == "" {
		if c.Request.Referer() != "" {
			return_to = c.Request.Referer()
		} else {
			return_to = "/"
		}
	}
	return c.Redirect("%s", return_to)
}

func (c *CommonDomain) SetLogin(account *models.Account) {
	c.Session[ACCOUNTID] = fmt.Sprint(account.Id)
	c.Session[EMAIL] = account.Email
}

func (c *CommonDomain) SetLogout() {
	delete(c.Session, ACCOUNTID)
	delete(c.Session, EMAIL)
}
