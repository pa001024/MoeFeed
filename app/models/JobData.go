package models

import (
	"encoding/json"
	"sort"
)

type KeyPair struct {
	Key   string
	Value interface{}
}
type JobData struct {
	Id        int64
	Name      string `qbs:"index,size:32,notnull"`
	Type      int32  `qbs:"notnull"`
	ProjectId int64  `qbs:"index,notnull"`
	Project   *Project
	Data      string
}

func (this *JobData) ViewData() (rst []KeyPair) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(this.Data), m)
	if err != nil {
		return nil
	}
	sl := make(sort.StringSlice, len(m))
	op := 0
	for v, _ := range m {
		sl[op] = v
		op++
	}
	sl.Sort()
	rst = make([]KeyPair, len(m))
	for i, v := range sl {
		rst[i] = KeyPair{v, m[v]}
	}
	return
}
