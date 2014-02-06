package controllers

import (
	r "github.com/robfig/revel"
)

type Organization struct{ PlatformDomain }

// 显示组织独立页
func (c Organization) Profile() r.Result {
	return c.Render()
}

func (c Organization) Show(user string) r.Result {
	return c.Render()
}
