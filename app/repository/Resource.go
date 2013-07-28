package repository

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
	"github.com/pa001024/MoeWorker/util"
)

var ResourceRepo *Resource

type Resource struct{}

func (this *Resource) Put(model *models.Resource) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

func (this *Resource) PutAndStone(resource *models.Resource, r multipart.File) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	defer r.Close()
	// 计算Hash
	m := md5.New()
	r.Seek(0, 0)
	io.Copy(m, r)
	resource.Hash = fmt.Sprintf("%x", m.Sum(nil))
	fn := resource.FileName()
	os.MkdirAll(fn[:len(fn)-32], 0644)
	f, err := os.Create(fn)
	if err != nil {
		util.Log(err)
		return
	}
	defer f.Close()
	r.Seek(0, 0)
	size, _ := io.Copy(f, r)
	resource.Size = size
	q.Save(resource)
}

func (this *Resource) GetFile(resource *models.Resource) io.ReadCloser {
	// TODO: 缓存?
	f, err := os.Open(resource.FileName())
	if err != nil {
		util.Log(err)
	}
	return f
}
func (this *Resource) Delete(model *models.Resource) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

// 主键
func (this *Resource) GetById(id int64) *models.Resource {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Resource{Id: id}
	q.Find(obj)
	if obj.Name == "" {
		return nil
	}
	return obj
}

// 联合聚集索引
func (this *Resource) GetByProjectAndName(name string, projectId int64) *models.Resource {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Resource{}
	q.Where("name = ? and project_id = ?", name, projectId).Find(obj)
	if obj.ProjectId == 0 {
		return nil
	}
	return obj
}

// 列出项目所有Resource
func (this *Resource) FindByProject(projectId int64) (obj []*models.Resource) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.WhereEqual("project_id", projectId).FindAll(&obj)
	if err != nil {
		return nil
	}
	return
}
