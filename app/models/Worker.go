package models

import (
	"bytes"
	"log"
	"time"

	"github.com/pa001024/MoeWorker/daemon"

// "github.com/pa001024/MoeWorker/filter"
// "github.com/pa001024/MoeWorker/source"
// "github.com/pa001024/MoeWorker/target"
// "github.com/pa001024/MoeWorker/util"
)

type Worker struct {
	Id         int64
	ConfigJson string
	ProjectId  int64 `qbs:"index,notnull"`
	Project    *Project
	Created    time.Time
	Updated    time.Time
}

func (this *Worker) JobConfig() *daemon.JobConfig {
	r := bytes.NewBufferString(this.ConfigJson)
	conf := &daemon.JobConfig{}
	if err := conf.Load(r); err != nil {
		log.Fatalln(err)
		return nil
	}
	return conf
}
