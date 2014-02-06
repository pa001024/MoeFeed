package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	r "github.com/robfig/revel"
)

// 项目控制器
type Project struct{ PlatformDomain }

///////////////////////////
// [动]具体动作 如增删改 //
///////////////////////////

// [动]创建项目
func (c Project) DoCreate(project *models.Project) r.Result {
	u, po := c.CheckUser()
	defer po.Close()
	if u == nil {
		c.Redirect("/login?return_to=/new")
	}
	project.OwnerId = u.Id
	c.Validation.Check(project.Name, r.Required{}, r.MinSize{2}, r.MaxSize{64}, r.Match{_accountRegex}).
		Message("project.invalid")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/new")
	}
	po.Put(project)
	r.INFO.Printf("%+v\n", project)

	return c.Redirect("/%v/%v", u.Name, project.Name)
}

// [动]重命名项目
func (c Project) Rename(user, project string) r.Result {
	c.CheckUserAndClose()
	return c.Render()
}

// [动]删除项目
func (c Project) DoDelete(user, project string) r.Result {
	c.CheckUserAndClose()
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
	u := c.CheckUserAndClose()
	return c.Redirect("/%s", u.Name)
}

// [静]创建项目前端
func (c Project) Create() r.Result {
	u := c.CheckUserAndClose()
	if u == nil {
		c.Redirect("/login?return_to=/new")
	}
	return c.Render()
}

// [静]删除项目
func (c Project) Delete(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mAdminable"].(bool) {
		return c.Forbidden(c.Message("project.edit.nopermission"))
	}
	return c.Render()
}

// [静]浏览页面
func (c Project) Explore() r.Result {
	c.CheckUserAndClose()
	return c.Render()
}

// [静]设置页面前端
func (c Project) Setting(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mAdminable"].(bool) {
		return c.Forbidden(c.Message("project.edit.nopermission"))
	}
	return c.Render()
}

// [静]显示项目独立页
func (c Project) Show(user, project string) r.Result {
	u, p, po := c.CheckAccessProjectRenderArgs(user, project)
	defer po.Close()
	if p == nil {
		return c.NotFound(c.Message("project.notfound"))
	}
	if u == nil && p.Type != models.ProjectPublic {
		return c.Forbidden(c.Message("project.view.nopermission"))
	}
	mSources := po.FindSource(p.Id)
	mFilters := po.FindFilter(p.Id)
	mTargets := po.FindTarget(p.Id)
	mResources := po.FindResource(p.Id)
	mCallbacks := po.FindCallback(p.Id)
	return c.Render(mSources, mFilters, mTargets, mResources, mCallbacks)
}
