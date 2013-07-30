package models

import ()

type SourceLog struct {
	Id       int64
	SourceId int64
	Source   *Source
	Title    string
	Content  string
}

func (this *SourceLog) Name() string {
	return this.Title
}
