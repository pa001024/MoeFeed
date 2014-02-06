package controllers

import (
	"fmt"
	"path"

	"github.com/pa001024/MoeFeed/app/models"
	r "github.com/robfig/revel"
)

// 资源控制器
type Resource struct{ PlatformDomain }

// [动][写]
func (c Resource) DoCreate(user, project string, resource *models.Resource) r.Result {
	u, p, po := c.CheckAccessProjectRenderArgs(user, project)
	defer po.Close()
	if u == nil {
		c.Flash.Error(c.Message("project.edit.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	if p == nil {
		return c.NotFound(c.Message("project.notfound"))
	}
	resource.ProjectId = p.Id
	fo, fh, err := c.Request.FormFile("file")
	if err == nil && fh.Filename != "" {
		if resource.Type == -1 {
			mime := fh.Header.Get("Content-Type")
			ext := path.Ext(fh.Filename)
			fmt.Println(mime, ext)
		}
	} else {
		c.Flash.Error("请上传文件: %v", err)
		return c.Redirect("/%s/%s/resource/new", user, project)
	}
	if po.GetResource(resource.Name, resource.ProjectId) != nil {
		c.Flash.Error("该资源名称已存在")
		return c.Redirect("/%s/%s/resource/new", user, project)
	}
	po.PutAndStoneResource(resource, fo)
	return c.Redirect("/%s/%s", user, project)
}

// [静]显示单个资源
func (c Resource) Show(user, project string) r.Result {
	c.CheckAccessProjectRenderArgsAndClose(user, project)
	if !c.RenderArgs["mReadable"].(bool) {
		c.Flash.Error(c.Message("project.view.nopermission"))
		return c.Redirect("/%s/%s", user, project)
	}
	return c.Render()
}

// [静]创建前端
func (c Resource) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
