package models

import (
	"fmt"
	"time"
)

// 项目
type Project struct {
	Id          int64
	Type        int16  `qbs:"notnull"`
	Name        string `qbs:"size:64,index,notnull"`
	DisplayName string `qbs:"size:64,index,notnull"`
	Desc        string `qbs:"size:140"`
	OwnerId     int64  `qbs:"index,notnull"`
	Owner       *PlatformUser
	Created     time.Time
	Updated     time.Time
}

// enum Project.Type
const (
	ProjectPublic = iota
	ProjectPrivate
)

// 返回字符串
func (this *Project) String() string {
	return fmt.Sprintf("Project:%s/%s", this.Owner.Account.Username, this.Name)
}

// 返回格式化的更新时间
func (this *Project) UpdateTime() string {
	return this.Updated.Format("1月2日 15:04")
}

func (this *Project) TypeName() string {
	if this.Type == ProjectPublic {
		return "project.public"
	} else {
		return "project.private"
	}
}

// 项目访问权限
type ProjectAccess struct {
	Id        int64
	Type      int16 `qbs:"notnull"`
	UserId    int64 `qbs:"notnull"`
	User      *PlatformUser
	ProjectId int64 `qbs:"notnull"`
	Project   *Project
}

// enum ProjectAccess.Type
const (
	AccessDeny      = iota // 完全拒绝权限
	AccessRead             // 只读权限
	AccessReadWrite        // 读写权限
	AccessAdmin            // 管理权限
	AccessOwner            // 所有者权限
)

func (this *ProjectAccess) Access() int16 {
	if this == nil {
		return AccessDeny
	} else {
		return this.Type
	}
}
