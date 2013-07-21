package models

import (
	"fmt"
	"math/rand"
)

type UserCode struct {
	Id     int64
	Code   string `qbs:"size:16,index,notnull"`
	UserId int64
	User   *User
}

func (this *UserCode) GenerateCode() {
	this.Code = fmt.Sprintf("%x", rand.Uint32())
}
