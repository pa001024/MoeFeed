package models

import (
	"time"
)

// 用户操作动态
type UserStatus struct {
	Id        int64
	Type      int16  `qbs:"notnull"`
	Desc      string `qbs:"size:140,notnull"`
	Link      string `qbs:"size:100,notnull"`
	UserId    int64  `qbs:"index,notnull"`
	User      *User
	ProjectId int64 `qbs:"index,notnull"`
	Project   *Project
	Created   time.Time
}

// enum UserStatus.Type
const (
	StatusCreateProject = iota
	StatusDelectProject
	StatusSendFeed
	StatusCreateIssues
)
