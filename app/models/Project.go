package models

import (
	"fmt"
	r "github.com/robfig/revel"
	"time"
)

// 项目
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

// 传递式验证
func (this *Project) Validate(v *r.Validation) {
	v.Check(this.Name, r.Required{}, r.MinSize{2}, r.MaxSize{64}, r.Match{projRegex})

	v.Required(this.OwnerId)
}

// 返回字符串
func (this *Project) String() string {
	return fmt.Sprintf("Project:%s/%s", this.Owner.Username, this.Name)
}

// 返回格式化的更新时间
func (this *Project) UpdateTime() string {
	return this.Updated.Format("1月2日 15:04")
}

// 是否可以被用户编辑
func (this *Project) CanEdit(user *User) bool {
	if user != nil && this.OwnerId == user.Id {
		return true
	}
	return false
}
