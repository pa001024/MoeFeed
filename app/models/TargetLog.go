package models

import ()

type TargetLog struct {
	Id       int64
	TargetId int64
	Target   *Target
	Title    string
	Content  string
}

func (this *TargetLog) Name() string {
	return this.Title
}
