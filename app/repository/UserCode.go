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

// 删除用户所有UserCode
func (this *UserCode) DeleteByOwner(ownerId int64) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.WhereEqual("owner_id", ownerId).Delete(new(UserCode))
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

func (this *UserCode) FindByCode(code string) (obj []*models.UserCode) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.WhereEqual("code", code).FindAll(&obj)
	if err != nil {
		return nil
	}
	return obj
}

// 列出用户所有UserCode
func (this *UserCode) FindByOwner(ownerId int64) (obj []*models.UserCode) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	err = q.WhereEqual("owner_id", ownerId).FindAll(&obj)
	if err != nil {
		return nil
	}
	return
}

// 精确获取用户
func (this *UserCode) GetByOwnerAndCode(code string, ownerId int64) (obj *models.UserCode) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj = &models.UserCode{}
	err = q.Where("code = ? and owner_id = ?", code, ownerId).Find(obj)
	if err != nil {
		return nil
	}
	return
}
