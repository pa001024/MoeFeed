package models

import (
	"bytes"
	"io"
	"os"

	"github.com/pa001024/MoeWorker/util"
)

// 资源
type Resource struct {
	Id        int64
	Name      string `qbs:"size:32,notnull"`
	Size      int32  `qbs:"notnull"`
	Hash      string `qbs:"size:32,notnull"` // 文件储存方式 /hash[:2]/hash
	ProjectId int64  `qbs:"index,notnull"`
	Project   *Project
}

func (this *Resource) FileName() string {
	return "data/" + this.Hash[:2] + "/" + this.Hash // TODO: 移动到配置文件
}

func (this *Resource) File() []byte {
	// TODO: 缓存?
	f, err := os.Open(this.FileName())
	if err != nil {
		util.DebugLog(err)
	}
	buf := &bytes.Buffer{}
	io.Copy(buf, f)
	return buf.Bytes()
}
