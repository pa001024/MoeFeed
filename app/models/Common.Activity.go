package models

import (
	"time"
)

// 用户操作日志
type UserActionRecord struct {
	Id      int64
	Action  string `qbs:"size:16,notnull"`
	Data    string `qbs:"size:256,notnull"`
	Created time.Time
}

// 用户登陆日志
type UserLoginRecord struct {
	Id      int64
	LoginIp string    `qbs:"size:32,notnull"` // 登录IP
	Created time.Time `qbs:"notnull"`
}
