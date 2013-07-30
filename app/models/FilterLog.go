package models

import ()

type FilterLog struct {
	Id       int64
	FilterId int64
	Filter   *Filter
	Title    string
	Content  string
}

func (this *FilterLog) Name() string {
	return this.Title
}
