package controllers

import (
	r "github.com/robfig/revel"
)

type Dashboard struct{ PlatformDomain }

func (c Dashboard) Index() r.Result {
	u, po := c.CheckUser()
	if u == nil {
		if po != nil {
			po.Close()
		}
		return c.Render()
	}
	mProjects := po.FindProjectByOwner(u.Id)
	if po != nil {
		po.Close()
	}
	return c.Render(mProjects)
}

func (c *Dashboard) Help() r.Result {
	c.CheckUserAndClose()
	return c.Render()
}
