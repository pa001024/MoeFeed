package controllers

import (
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

type Api struct {
	App
}

func (c Api) Index() r.Result {
	return c.Render()
}

func (c Api) Test() r.Result {
	return c.NotFound("恩")
}
