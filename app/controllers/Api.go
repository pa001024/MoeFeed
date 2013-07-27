package controllers

import (
	// repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 外部API
type Api struct{ App }

func (c Api) Index() r.Result {
	return c.Render()
}
