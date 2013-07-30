package controllers

import (
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

type Dashboard struct{ App }

func (c Dashboard) Index() r.Result {
	u := c.CheckUser()
	mProjects := repo.ProjectRepo.FindByOwner(u.Id)
	return c.Render(mProjects)
}
