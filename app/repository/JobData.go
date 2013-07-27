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
