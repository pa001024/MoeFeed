package models

import (
	"time"
)

// 路线
type Channel struct {
	Id        int64
	FromId    int64 `qbs:"notnull"`
	From      *Source
	ToId      int64 `qbs:"notnull"`
	To        *Target
	Status    int16
	ProjectId int64 `qbs:"index,notnull"`
	Project   *Project
	Created   time.Time
	Updated   time.Time
}

// enum Channel.Status
const (
	ChannelEnable  = iota // 启用
	ChannelDisable        // 禁用
)

// 获取名字
func (this *Channel) Name() string {
	return this.From.Name + "->" + this.To.Name
}
