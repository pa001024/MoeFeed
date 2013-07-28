package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var UserStatusRepo *UserStatus

type UserStatus struct{}

func (this *UserStatus) Put(model *models.UserStatus) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

func (this *UserStatus) Delete(model *models.UserStatus) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

// 主键
func (this *UserStatus) GetById(id int64) *models.UserStatus {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.UserStatus{Id: id}
	q.Find(obj)
	if obj.ProjectId == 0 {
		return nil
	}
	return obj
}

// 列出用户所有UserStatus
func (this *UserStatus) FindByUser(userId int64) (obj []*models.UserStatus) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.WhereEqual("user_id", userId).FindAll(&obj)
	if err != nil {
		return nil
	}
	return
}

// 列出项目所有UserStatus
func (this *UserStatus) FindByProject(projectId int64) (obj []*models.UserStatus) {
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

// 列出用户项目所有UserStatus
func (this *UserStatus) FindByUserAndProject(userId, projectId int64) (obj []*models.UserStatus) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.Where("user_id = ? and project_id = ?", userId, projectId).FindAll(&obj)
	if err != nil {
		return nil
	}
	return
}
