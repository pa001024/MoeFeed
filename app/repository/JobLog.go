package repository

import (
	"github.com/coocood/qbs" // TODO: NOSQL
	"github.com/pa001024/MoeFeed/app/models"
)

var JobLogRepo *JobLog

type JobLog struct{}

func (this *JobLog) Put(model *models.JobLog) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

func (this *JobLog) Delete(model *models.JobLog) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

// 主键
func (this *JobLog) GetById(id int64) *models.JobLog {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.JobLog{Id: id}
	q.Find(obj)
	if obj.ChannelId == 0 {
		return nil
	}
	return obj
}

// 列出项目所有JobLog
func (this *JobLog) FindByChannel(channelId int64) (obj []*models.JobLog) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.WhereEqual("channel_id", channelId).FindAll(&obj)
	if err != nil {
		return nil
	}
	return
}
