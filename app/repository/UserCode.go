package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var UserCodeRepo *UserCode

type UserCode struct{}

func (this *UserCode) Put(model *models.UserCode) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(model)
}

// 删除
func (this *UserCode) Delete(model *models.UserCode) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

func (this *UserCode) GetById(id int64) *models.UserCode {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.UserCode{Id: id}
	q.Find(obj)
	if obj.Code == "" {
		return nil
	}
	return obj
}

func (this *UserCode) GetByCode(code string) *models.UserCode {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.UserCode{}
	q.WhereEqual("code", code).Find(obj)
	if obj.Id == 0 {
		return nil
	}
	return obj
}

// 列出用户所有项目
func (this *UserCode) FindByOwner(ownerId int64) (obj []*models.UserCode) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.WhereEqual("owner_id", ownerId).OrderByDesc("updated").FindAll(&obj)
	if err != nil {
		return nil
	}
	return
}
