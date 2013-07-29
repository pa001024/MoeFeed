package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	"fmt"
)

// 项目控制器
type Project struct{ App }

// 统一参数解析
func (c Project) CheckProject(user, project string) *models.Project {
	if vp := c.RenderArgs["mProject"]; vp != nil {
		return vp.(*models.Project)
	}
	mProject := repo.ProjectRepo.GetByName(user, project)
	if mProject != nil {
		c.RenderArgs["mProject"] = mProject
		return mProject
	}
	return nil
}

// 检查编辑权限
func (c Project) CheckOwnerProject(user, project string) (*models.User, *models.Project) {
	u := c.CheckUser()
	p := c.CheckProject(user, project)
	if u.Id == p.OwnerId {
		return u, p
	}
	return nil, p
}

///////////////////////////
// [动]具体动作 如增删改 //
///////////////////////////

// [动]创建项目
func (c Project) DoCreate(project *models.Project) r.Result {
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
func (c Project) Rename(user, project string) r.Result {
	c.CheckUser()
	return c.Render()
}

// [动]删除项目
func (c Project) DoDelete(user, project string) r.Result {
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

// [静]删除项目
func (c Project) Delete(user, project string) r.Result {
	u, _ := c.CheckOwnerProject(user, project)
	if u == nil {
		c.Flash.Error("你没有权限编辑该项目")
		return c.Redirect("/%s/%s", user, project)
	}
	return c.Render()
}

// [静]浏览页面
func (c Project) Explore() r.Result {
	c.CheckUser()
	return c.Render()
}

// [静]设置页面前端
func (c Project) Setting(user, project string) r.Result {
	c.CheckUser()
	return c.Render()
}

// [静]显示项目独立页
func (c Project) Show(user, project string) r.Result {
	u := c.CheckUser()
	p := c.CheckProject(user, project)
	if p == nil {
		return c.NotFound("没有此项目")
	}
	mSources := repo.SourceRepo.FindByProject(p.Id)
	mFilters := repo.FilterRepo.FindByProject(p.Id)
	mTargets := repo.TargetRepo.FindByProject(p.Id)
	mResources := repo.ResourceRepo.FindByProject(p.Id)
	mCallbacks := repo.CallbackRepo.FindByProject(p.Id)
	mEditable := false
	if u != nil {
		mEditable = u.Id == p.OwnerId
	}
	return c.Render(mEditable, mSources, mFilters, mTargets, mResources, mCallbacks)
}
