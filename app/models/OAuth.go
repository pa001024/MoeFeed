package models

import (
	"time"
)

type OAuth struct {
	Code    string `qbs:"size:24,unique"`
	UserId  int64
	User    *User
	Expires time.Time
}
