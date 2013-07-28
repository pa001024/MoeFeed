package controllers

import (
	"fmt"
	"path"

	"github.com/pa001024/MoeFeed/app/models"
	repo "github.com/pa001024/MoeFeed/app/repository"
	r "github.com/robfig/revel"
)

// 资源控制器
type Resource struct{ Project }

// [动][写]
func (c Resource) PostCreate(user, project string, resource *models.Resource) r.Result {
	u := c.CheckUser()
	p := c.CheckProject(user, project)
	if p == nil {
		return c.NotFound("找不到该项目")
	}
	if u == nil {
		c.Flash.Error("请先登录")
		return c.Redirect("/%s/%s/resource/new", user, project)
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
	if repo.ResourceRepo.GetByProjectAndName(resource.Name, resource.ProjectId) != nil {
		c.Flash.Error("该资源名称已存在")
		return c.Redirect("/%s/%s/resource/new", user, project)
	}
	repo.ResourceRepo.PutAndStone(resource, fo)
	return c.Redirect("/%s/%s", user, project)
}

// [静]显示单个资源
func (c Resource) Show(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}

// [静]创建前端
func (c Resource) Create(user, project string) r.Result {
	c.CheckUser()
	c.CheckProject(user, project)
	return c.Render()
}
