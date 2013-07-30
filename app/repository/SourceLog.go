package repository

import (
	"github.com/coocood/qbs" // TODO: NOSQL
	"github.com/pa001024/MoeFeed/app/models"
)

var SourceLogRepo *SourceLog

type SourceLog struct{}

func (this *SourceLog) Put(model *models.SourceLog) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

func (this *SourceLog) Delete(model *models.SourceLog) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

// 主键
func (this *SourceLog) GetById(id int64) *models.SourceLog {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.SourceLog{Id: id}
	q.Find(obj)
	if obj.Title == "" {
		return nil
	}
	return obj
}

// 列出项目所有SourceLog
func (this *SourceLog) FindBySource(sourceId int64) (obj []*models.SourceLog) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.WhereEqual("source_id", sourceId).FindAll(&obj)
	if err != nil {
		return nil
	}
	return
}
