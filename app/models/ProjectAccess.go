package models

type ProjectAccess struct {
	Id        int64
	Type      int16 `qbs:"notnull"`
	UserId    int64 `qbs:"index,notnull"`
	User      *User
	ProjectId int64 `qbs:"index,notnull"`
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
