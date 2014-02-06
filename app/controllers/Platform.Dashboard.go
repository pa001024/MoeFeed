package controllers

import (
	r "github.com/robfig/revel"
)

type Dashboard struct{ PlatformDomain }

func (c Dashboard) Index() r.Result {
	u, po := c.CheckUser()
	defer po.Close()
	if u == nil {
		return c.Render()
	}
	mProjects := po.FindProjectByOwner(u.Id)
	return c.Render(mProjects)
}
