package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"

	"fmt"
)

// 项目控制器
type Project struct{ App }

// 统一参数解析 mProject setter
func (c *Project) CheckProject(user, project string) *models.Project {
	if vp := c.RenderArgs["mProject"]; vp != nil {
		return vp.(*models.Project)
	}
	mProject := repo.ProjectRepo.GetByName(user, project)
	c.RenderArgs["mProject"] = mProject
	return mProject
}

// 检查权限
func (c *Project) CheckAccessProject(user, project string, access int16) (*models.User, *models.Project, int16) {
	u, p := c.CheckUser(), c.CheckProject(user, project)
	if p == nil {
		return nil, nil, models.AccessDeny
	}
	if u != nil {
		if u.Id == p.OwnerId {
			return u, p, models.AccessOwner
		}
		access = repo.ProjectRepo.GetAccess(u.Id, p.Id).Access()
	}
	if access < models.AccessRead && p.Type == models.ProjectPublic {
		access = models.AccessRead
	}
	return u, p, access
}

// 权限渲染
func (c *Project) CheckAccessRenderArgs(access int16) {
	c.RenderArgs["mReadable"] = access >= models.AccessRead
	c.RenderArgs["mEditable"] = access >= models.AccessReadWrite
	c.RenderArgs["mAdminable"] = access >= models.AccessAdmin
	c.RenderArgs["mOwnerable"] = access >= models.AccessOwner
}

// 检查编辑权限 mReadable setter
func (c *Project) CheckReadableProject(user, project string) (*models.User, *models.Project) {
	u, p, a := c.CheckAccessProject(user, project, models.AccessRead)
	c.CheckAccessRenderArgs(a)
	return u, p
}

// 检查编辑权限 mEditable setter
func (c *Project) CheckEditableProject(user, project string) (*models.User, *models.Project) {
	u, p, a := c.CheckAccessProject(user, project, models.AccessReadWrite)
	c.CheckAccessRenderArgs(a)
	return u, p
}

// 检查管理权限 mAdminable setter
func (c *Project) CheckAdminableProject(user, project string) (*models.User, *models.Project) {
	u, p, a := c.CheckAccessProject(user, project, models.AccessAdmin)
	c.CheckAccessRenderArgs(a)
	return u, p
}

///////////////////////////
// [动]具体动作 如增删改 //
///////////////////////////

// [动]创建项目
func (c Project) DoCreate(project *models.Project) r.Result {
	u := c.CheckUser()
	if u == nil {
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
	u := c.CheckUser()
	if u == nil {
		c.Redirect("/login?return_to=/new")
	}
	return c.Render()
}

// [静]删除项目
func (c Project) Delete(user, project string) r.Result {
	u, _ := c.CheckAdminableProject(user, project)
	if u == nil {
		return c.Forbidden("你没有权限编辑该项目")
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
	u, _ := c.CheckAdminableProject(user, project)
	if u == nil {
		return c.Forbidden("你没有权限编辑该项目")
	}
	return c.Render()
}

// [静]显示项目独立页
func (c Project) Show(user, project string) r.Result {
	u, p := c.CheckReadableProject(user, project)
	if p == nil {
		return c.NotFound("没有此项目")
	}
	if u == nil && p.Type != models.ProjectPublic {
		return c.Forbidden("你没有权限查看该项目")
	}
	mSources := repo.SourceRepo.FindByProject(p.Id)
	mFilters := repo.FilterRepo.FindByProject(p.Id)
	mTargets := repo.TargetRepo.FindByProject(p.Id)
	mResources := repo.ResourceRepo.FindByProject(p.Id)
	mCallbacks := repo.CallbackRepo.FindByProject(p.Id)
	return c.Render(mSources, mFilters, mTargets, mResources, mCallbacks)
}
