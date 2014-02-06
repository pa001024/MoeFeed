package models

import (
	r "github.com/robfig/revel"
)

var (
	_accountSecret = "mOeFeEd"
)

func PostInit() {
	_accountSecret = r.Config.StringDefault("salt.user", "mOeFeEd")
}
