package controllers

import (
	r "github.com/robfig/revel"
)

type OAuth struct {
	*r.Controller
}

func (c OAuth) Index() r.Result {
	return c.Render()
}
