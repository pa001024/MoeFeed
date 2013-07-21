package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"
)

var (
	ProjectRepo *Project
)

type Project struct{}

func (this *Project) Put(project *models.Project) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(project)
}

// 因为Project没有聚集索引 所以直接使用这个
func (this *Project) GetById(id int64) *models.Project {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.Project{Id: id}
	q.Find(obj)
	if obj.Name == "" {
		return nil
	}
	return obj
}

func (this *Project) GetByName(userName, projectName string) *models.Project {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	u := &models.User{}
	q.WhereEqual("username", userName).Find(u)
	if u.Id == 0 {
		return nil
	}
	obj := &models.Project{}
	q.Where("owner_id = ? and name = ?", u.Id, projectName).Find(obj)
	if obj.Id == 0 {
		return nil
	}
	return obj
}

// 列出用户所有项目
func (this *Project) FindByOwner(ownerId int64) (obj []*models.Project) {
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
