package repository

import (
	"strings"

	"github.com/pa001024/MoeFeed/app/models"
)

func CommonRepo(q ...*QbsRepository) *Common {
	if len(q) > 0 {
		return &Common{q[0]}
	}
	return &Common{QbsRepo()}
}

type Common struct{ *QbsRepository }

// 聚集索引
func (this *Common) GetAccount(id interface{}) (m *models.Account) {
	this.GetRef(&m, "account.id", id)
	return
}

// 聚集索引
func (this *Common) GetAccountByName(username string) (m *models.Account) {
	this.GetRef(&m, "username", username)
	return
}

// 聚集索引
func (this *Common) GetAccountByEmail(email string) (m *models.Account) {
	this.GetRef(&m, "email", email)
	return
}

// 自动获取
func (this *Common) GetAccountByNameOrEmail(nameOrEmail string) *models.Account {
	if strings.ContainsRune(nameOrEmail, '@') {
		return this.GetAccountByEmail(nameOrEmail)
	} else {
		return this.GetAccountByName(nameOrEmail)
	}
}

// 获取Email验证
func (this *Common) GenerateEmailVerfiy(account_id int64) (m *models.AccountEmailVerify) {
	m = &models.AccountEmailVerify{AccountId: account_id}
	m.GenerateCode()
	this.Put(m)
	return
}

// 验证Email验证
func (this *Common) GetEmailVerfiy(code string, username string) (m *models.AccountEmailVerify) {
	this.GetNRef(&m, "code=?", code, "account.username=?", username)
	return
}
