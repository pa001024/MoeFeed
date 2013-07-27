package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var ChannelRepo *Channel

type Channel struct{}

func (this *Channel) Put(model *models.Channel) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

// 主键
func (this *Channel) GetById(id int64) *models.Channel {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Channel{Id: id}
	q.Find(obj)
	if obj.ProjectId == 0 {
		return nil
	}
	return obj
}

// 列出项目所有Channel
func (this *Channel) FindByProject(projectId int64) (obj []*models.Channel) {
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
