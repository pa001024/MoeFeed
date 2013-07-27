package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var CallbackRepo *Callback

type Callback struct{}

func (this *Callback) Put(model *models.Callback) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

// 主键
func (this *Callback) GetById(id int64) *models.Callback {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Callback{Id: id}
	q.Find(obj)
	if obj.ProjectId == 0 {
		return nil
	}
	return obj
}

// 列出项目所有Callback
func (this *Callback) FindByProject(projectId int64) (obj []*models.Callback) {
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
