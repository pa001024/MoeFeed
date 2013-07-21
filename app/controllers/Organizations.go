package controllers

import (
	r "github.com/robfig/revel"
)

type Organizations struct {
	*r.Controller
}

// 显示组织独立页
func (c Organizations) Profile() r.Result {
	return c.Render()
}

func (c Organizations) Show(user string) r.Result {
	return c.Render()
}
