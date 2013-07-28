package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var SourceRepo *Source

type Source struct{}

func (this *Source) Put(model *models.Source) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

func (this *Source) Delete(model *models.Source) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

// 主键
func (this *Source) GetById(id int64) *models.Source {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Source{Id: id}
	q.Find(obj)
	if obj.Name == "" {
		return nil
	}
	return obj
}

// 联合聚集索引
func (this *Source) GetByProjectAndName(name string, projectId int64) *models.Source {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Source{}
	q.Where("name = ? and project_id = ?", name, projectId).Find(obj)
	if obj.ProjectId == 0 {
		return nil
	}
	return obj
}

// 列出项目所有Source
func (this *Source) FindByProject(projectId int64) (obj []*models.Source) {
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
