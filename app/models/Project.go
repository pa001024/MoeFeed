package models

import (
	"fmt"
	r "github.com/robfig/revel"
	"time"
)

type Project struct {
	Id          int64
	Name        string `qbs:"size:64,index,notnull"`
	DisplayName string `qbs:"size:64,index,notnull"`
	Desc        string `qbs:"size:140"`
	OwnerId     int64  `qbs:"index,notnull"`
	Owner       *User
	Created     time.Time
	Updated     time.Time
}

var (
	projRegex = userRegex
)

func (this *Project) Validate(v *r.Validation) {
	v.Check(this.Name, r.Required{}, r.MinSize{2}, r.MaxSize{64}, r.Match{projRegex})

	v.Required(this.OwnerId)
}

func (this *Project) String() string {
	return fmt.Sprintf("Project:%s/%s", this.Owner.Username, this.Name)
}

func (this *Project) UpdateTime() string {
	return this.Updated.Format("1月2日 15:04")
}
