package models

import ()

type JobLog struct {
	Id        int64
	ChannelId int64
	Channel   *Channel
	Title     string
	Content   string
}

func (this *JobLog) Name() string {
	return this.Title
}
