package controllers

import (
	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
)

// 平台
type PlatformDomain struct{ CommonDomain }

// 用户状态持久化
func (c *PlatformDomain) CheckUser() (u *models.PlatformUser, po *repo.Platform) {
	po = repo.PlatformRepo()
	if id, ok := c.Session[ACCOUNT]; ok {
		u = po.GetUser(id)
		c.RenderArgs["mUser"] = u
	}
	return
}

// 用户状态持久化
func (c *PlatformDomain) CheckUserAndClose() *models.PlatformUser {
	u, po := c.CheckUser()
	po.Close()
	return u
}

// 统一参数解析 mProject setter
func (c *PlatformDomain) CheckProject(user, project string) (mProject *models.Project, po *repo.Platform) {
	po = repo.PlatformRepo()
	mProject = po.GetProject(user, project)
	c.RenderArgs["mProject"] = mProject
	return
}

// 统一参数解析 mProject setter
func (c *PlatformDomain) CheckProjectAndClose(user, project string) *models.Project {
	m, po := c.CheckProject(user, project)
	po.Close()
	return m
}

// 检查权限
func (c *PlatformDomain) CheckAccessProject(user, project string) (u *models.PlatformUser, p *models.Project, access int16, po *repo.Platform) {
	u = c.CheckUserAndClose()
	p, po = c.CheckProject(user, project)
	if p == nil {
		return
	}
	if u != nil {
		if u.Id == p.OwnerId {
			access = models.AccessOwner
			return
		}
		access = po.GetAccess(u.Id, p.Id).Access()
	}
	if access < models.AccessRead && p.Type == models.ProjectPublic {
		access = models.AccessRead
	}
	return
}

// 检查权限
func (c *PlatformDomain) CheckAccessProjectAndClose(user, project string) (*models.PlatformUser, *models.Project, int16) {
	u, p, a, po := c.CheckAccessProject(user, project)
	po.Close()
	return u, p, a
}

// 检查编辑权限 setter
func (c *PlatformDomain) CheckAccessProjectRenderArgs(user, project string) (*models.PlatformUser, *models.Project, *repo.Platform) {
	u, p, a, po := c.CheckAccessProject(user, project)
	c.RenderArgs["mReadable"] = a >= models.AccessRead
	c.RenderArgs["mEditable"] = a >= models.AccessReadWrite
	c.RenderArgs["mAdminable"] = a >= models.AccessAdmin
	c.RenderArgs["mOwnerable"] = a >= models.AccessOwner
	return u, p, po
}

// 检查编辑权限 setter
func (c *PlatformDomain) CheckAccessProjectRenderArgsAndClose(user, project string) (*models.PlatformUser, *models.Project) {
	u, p, po := c.CheckAccessProjectRenderArgs(user, project)
	po.Close()
	return u, p
}
