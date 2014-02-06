package models

import (
	"time"
)

// 用户动态
type UserActivity struct {
	Id        int64
	Link      string `qbs:"size:100,notnull"`
	UserId    int64  `qbs:"notnull"`
	User      *PlatformUser
	ProjectId int64 `qbs:"notnull"`
	Project   *Project
	Type      int16  `qbs:"notnull"`
	Title     string `qbs:"size:64,notnull"`
	Content   string `qbs:"size:256,notnull"`
	Created   time.Time
}

// enum UserActivity.Type
const (
	ActivityCreateProject = iota
	ActivityDelectProject
	ActivitySendFeed
	ActivityCreateIssues
)
