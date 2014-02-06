package models

import (
	"time"
)

type AdminUser struct {
	Id        int64
	AccountId int64 `qbs:"notnull"`
	Account   *Account
	Level     int16 `qbs:"notnull"`
	Created   time.Time
	Updated   time.Time
}

// enum AdminUser.Level
const (
	AdminLevelRead = iota
	AdminLevelReadWrite
	AdminLevelFull
)
