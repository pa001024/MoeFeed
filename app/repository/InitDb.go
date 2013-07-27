package repository

import (
	"github.com/coocood/qbs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pa001024/MoeFeed/app/models"
)

func init() {
	RegisterDb()
}

// 初始化数据库
func RegisterDb() {
	qbs.Register("mysql", DBSPEC, DBNAME, qbs.NewMysql())
	// 创建数据库
	m, err := qbs.GetMigration()
	if err != nil {
		panic(err)
	}
	defer m.Close()

	m.CreateTableIfNotExists(new(models.User))
	m.CreateTableIfNotExists(new(models.Project))
	m.CreateTableIfNotExists(new(models.UserCode))
	m.CreateTableIfNotExists(new(models.UserStatus))
	m.CreateTableIfNotExists(new(models.OAuth))
	m.CreateTableIfNotExists(new(models.Source))
	m.CreateTableIfNotExists(new(models.Target))
	m.CreateTableIfNotExists(new(models.Filter))
	m.CreateTableIfNotExists(new(models.Resource))
	m.CreateTableIfNotExists(new(models.Channel))
	m.CreateTableIfNotExists(new(models.Callback))

	// TODO: 日志表需要NOSQL
}
