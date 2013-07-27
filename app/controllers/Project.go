package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	"fmt"
)

// 项目控制器
type Project struct{ App }

///////////////////////////
// [动]具体动作 如增删改 //
///////////////////////////

// [动]创建项目
func (c Project) PostCreate(project *models.Project) r.Result {
	u := c.CheckUser()
	if u.Id == 0 {
		c.Redirect("/login?return_to=/new")
	}
	project.OwnerId = u.Id
	project.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/new")
	}
	repo.ProjectRepo.Put(project)
	fmt.Printf("%#v\n", project)

	return c.Redirect(fmt.Sprintf("/%v/%v", u.Username, project.Name))
}

// [动]重命名项目
func (c Project) Rename() r.Result {
	c.CheckUser()
	return c.Render()
}

// [动]删除项目
func (c Project) Delete() r.Result {
	c.CheckUser()
	return c.Render()
}

//////////////////
// [静]静态页面 //
//////////////////

// [链]创建项目同义词
func (c Project) CreateLink() r.Result {
	return c.Redirect("/new")
}

// [链]列表同义词
func (c Project) ListLink() r.Result {
	u := c.CheckUser()
	return c.Redirect("/" + u.Username)
}

// [静]创建项目前端
func (c Project) Create() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静]帮助页面
func (c Project) Help() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静]设置页面前端
func (c Project) Setting() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静]显示项目独立页
func (c Project) Show(user, project string) r.Result {
	mUser := c.CheckUser()
	mProject := repo.ProjectRepo.GetByName(user, project)
	if mProject == nil {
		return c.NotFound("项目不存在")
	}
	// mChannel := repo.ChannelRepo.FindByProject(mProject.Id)
	mSource := repo.SourceRepo.FindByProject(mProject.Id)
	mFilter := repo.FilterRepo.FindByProject(mProject.Id)
	mTarget := repo.TargetRepo.FindByProject(mProject.Id)
	mResource := repo.ResourceRepo.FindByProject(mProject.Id)
	mCallback := repo.CallbackRepo.FindByProject(mProject.Id)
	if mUser != nil {
		mEditable := mUser.Id == mProject.OwnerId

		return c.Render(mEditable, mUser, mProject, mSource, mFilter, mTarget, mResource, mCallback)
	}
	return c.Render(mProject)
}
