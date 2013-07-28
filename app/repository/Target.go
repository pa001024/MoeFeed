package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var TargetRepo *Target

type Target struct{}

func (this *Target) Put(model *models.Target) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

func (this *Target) Delete(model *models.Target) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

// 主键
func (this *Target) GetById(id int64) *models.Target {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Target{Id: id}
	q.Find(obj)
	if obj.Name == "" {
		return nil
	}
	return obj
}

// 联合聚集索引
func (this *Target) GetByProjectAndName(name string, projectId int64) *models.Target {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Target{}
	q.Where("name = ? and project_id = ?", name, projectId).Find(obj)
	if obj.ProjectId == 0 {
		return nil
	}
	return obj
}

// 列出项目所有Target
func (this *Target) FindByProject(projectId int64) (obj []*models.Target) {
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
