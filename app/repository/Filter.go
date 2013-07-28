package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var FilterRepo *Filter

type Filter struct{}

func (this *Filter) Put(model *models.Filter) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

// 主键
func (this *Filter) GetById(id int64) *models.Filter {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Filter{Id: id}
	q.Find(obj)
	if obj.Name == "" {
		return nil
	}
	return obj
}

// 联合聚集索引
func (this *Filter) GetByProjectAndName(name string, projectId int64) *models.Filter {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Filter{}
	q.Where("name = ? and project_id = ?", name, projectId).Find(obj)
	if obj.ProjectId == 0 {
		return nil
	}
	return obj
}

// 列出项目所有Filter
func (this *Filter) FindByProject(projectId int64) (obj []*models.Filter) {
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
