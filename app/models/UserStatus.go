package models

import (
	"strings"
	"time"

	"fmt"
	"github.com/pa001024/MoeWorker/util"
)

// 用户操作动态
type UserStatus struct {
	Id        int64
	Type      int8   `qbs:"notnull"`
	Desc      string `qbs:"size:140,notnull"`
	Link      string `qbs:"size:100,notnull"`
	UserId    int64  `qbs:"index,notnull"`
	User      *User
	ProjectId int64 `qbs:"index,notnull"`
	Project   *Project
	Created   time.Time
}

// TODO: 分离到VM层
type UserStatusVM struct {
	UserStatus

	// 瞬态
	AvatarUrl string `qbs:"omit"`
}

func NewUserStatusVM(org UserStatus) (rst *UserStatusVM) {
	rst = &UserStatusVM{UserStatus: org}
	rst.AvatarUrl = fmt.Sprintf("http://www.gravatar.com/avatar/%v?s=32&d=retro", util.Md5String(strings.ToLower(rst.User.AvatarEmail)))
	return
}

// enum UserStatus.Type
const (
	CreateProject = iota
	DelectProject
	SendFeed
	FoundIssues
)
