package controllers

import (
	repo "github.com/pa001024/MoeFeed/app/repository"

	// "github.com/pa001024/MoeFeed/app/models"
	// "github.com/pa001024/MoeFeed/app/service"
	// "github.com/pa001024/MoeWorker/util"
	r "github.com/robfig/revel"

// "log"
// "net/url"
// "strings"
)

type Dashboard struct{ App }

func (c Dashboard) Index() r.Result {
	u := c.CheckUser()
	projects := repo.ProjectRepo.FindByOwner(u.Id)
	return c.Render(projects)
}
