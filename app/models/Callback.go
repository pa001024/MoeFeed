package models

import (
	r "github.com/robfig/revel"
	"time"
)

// 路线
type Callback struct {
	Id        int64
	Name      string `qbs:"size:32,notnull"`
	Url       string `qbs:"size:32,unique,notnull"`
	Type      int32  `qbs:"notnull"`
	ProjectId int64  `qbs:"index,notnull"`
	Project   *Project
	Created   time.Time
	Updated   time.Time
}

func (this *Callback) Validate(v *r.Validation, password string) {
	v.Check(this.Name, r.Required{}, r.MinSize{2}, r.MaxSize{32})
	v.Check(this.Url, r.Required{}, r.MinSize{2}, r.MaxSize{32})
}
