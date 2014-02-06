package models

// 资源
type Resource struct {
	Id        int64
	Name      string `qbs:"size:32,notnull"`
	Type      int16  `qbs:"notnull"`
	Size      int64  `qbs:"notnull"`
	Hash      string `qbs:"size:32,index,notnull"` // 文件储存方式 /hash[:2]/hash
	ProjectId int64  `qbs:"notnull"`
	Project   *Project
}

// enum Resource.Type
const (
	ResourceImage  = iota // 图片文件
	ResourceText          // 文本文件
	ResourceScript        // 脚本文件
	ResourceOther         // 其他文件
)

func (this *Resource) FileName() string {
	return "./data/" + this.Hash[:2] + "/" + this.Hash // TODO:可配置化
}

func (this *Resource) ViewData() (rst []KeyPair) {
	rst = []KeyPair{
		{"文件名", this.Name},
		{"大小", this.Size},
		{"MD5", this.Hash},
	}
	return
}
