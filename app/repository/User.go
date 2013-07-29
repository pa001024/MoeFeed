package repository

import (
	"github.com/coocood/qbs"
	"github.com/pa001024/MoeFeed/app/models"

	"strings"
)

var UserRepo *User

type User struct{}

// 置入
func (this *User) Put(user *models.User) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Save(user)
}

// 删除
func (this *User) Delete(model *models.User) {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	q.Delete(model)
}

// 通用
func (this *User) GetBy(key string, value interface{}) *models.User {
	//////////////////
	q, err := qbs.GetQbs()
	assetsError(err)
	defer q.Close()
	//////////////////
	obj := &models.User{}
	q.WhereEqual(key, value).Find(obj)
	if obj.Username != "" {
		return obj
	}
	return nil
}

// 聚集索引 interface{}是为了兼容int64和string
func (this *User) GetById(id interface{}) *models.User {
	return this.GetBy("id", id)
}

// 聚集索引
func (this *User) GetByName(username string) *models.User {
	return this.GetBy("username", username)
}

// 聚集索引
func (this *User) GetByEmail(email string) *models.User {
	return this.GetBy("email", email)
}

func (this *User) GetByNameOrEmail(nameOrEmail string) *models.User {
	if strings.ContainsRune(nameOrEmail, '@') {
		return this.GetByEmail(nameOrEmail)
	} else {
		return this.GetByName(nameOrEmail)
	}
}

func assetsError(err error) {
	if err != nil {
		panic(err)
	}
}
