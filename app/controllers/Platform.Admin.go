package controllers

import (
	r "github.com/robfig/revel"
)

// 平台管理
type PlatformAdmin struct{ AdminDomain }

func (c *PlatformAdmin) Index() r.Result {
	u, po := c.CheckAdmin()
	po.Close()
	if u == nil {
		return c.Redirect("/")
	}
	return c.Render()
}
