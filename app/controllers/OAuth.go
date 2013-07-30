package controllers

import (
	r "github.com/robfig/revel"
)

type OAuth struct{ App }

func (c OAuth) Index() r.Result {
	return c.Render()
}
