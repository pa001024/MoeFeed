package models

import (
	"time"
)

// 邮件计划任务
type MailCron struct {
	Id       int64
	Touid    int64     `qbs:"notnull"`
	Email    string    `qbs:"size:100,notnull"`
	Sendtime time.Time `qbs:"notnull,index"`
}

// 邮件计划队列
type MailQueue struct {
	Id         int64
	MailCronId int64 `qbs:"notnull"`
	MailCron   *MailCron
}

type MyTask struct {
	Id         int64
	ActionType int32 `qbs:"notnull"`
}
