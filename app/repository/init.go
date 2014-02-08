package repository

import (
	"github.com/coocood/qbs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pa001024/MoeFeed/app/models"
	r "github.com/robfig/revel"
)

func init() {
	RegisterDb()
	// if r.DevMode {
	InitDb()
	// }
}

// 注册数据库
func RegisterDb() {
	_ = r.Config
	db_spec := "moefeed:moefeed@/moefeed?charset=utf8&parseTime=true&loc=Local" // r.Config.StringDefault("db.spec", "")
	db_name := "moefeed"                                                        // r.Config.StringDefault("db.name", "")
	qbs.Register("mysql", db_spec, db_name, qbs.NewMysql())
}

// 初始化数据库
func InitDb() {
	m, err := qbs.GetMigration()
	if err != nil {
		panic(err)
	}
	defer m.Close()
	// Common
	m.CreateTableIfNotExists(new(models.Account))
	m.CreateTableIfNotExists(new(models.AccountEmailVerify))
	// Platform
	m.CreateTableIfNotExists(new(models.UserActivity))
	m.CreateTableIfNotExists(new(models.PlatformUser))
	m.CreateTableIfNotExists(new(models.Project))
	m.CreateTableIfNotExists(new(models.ProjectAccess))
	m.CreateTableIfNotExists(new(models.Source))
	m.CreateTableIfNotExists(new(models.Filter))
	m.CreateTableIfNotExists(new(models.Target))
	m.CreateTableIfNotExists(new(models.Resource))
	m.CreateTableIfNotExists(new(models.Channel))
	m.CreateTableIfNotExists(new(models.Callback))
}
