package models

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
)

// 登录账号
type Account struct {
	Id       int64
	Username string `qbs:"size:32,unique,notnull"`
	Password string `qbs:"size:80,notnull"`
	Email    string `qbs:"size:100,unique,notnull"`
	Enable   bool   `qbs:"notnull"`
	Created  time.Time
	Updated  time.Time
}

// 验证密码
func (this *Account) ValidatePassword(password string) bool {
	bin := []byte(_accountSecret + password + this.Username)
	hbin, _ := base64.StdEncoding.DecodeString(this.Password)
	err := bcrypt.CompareHashAndPassword(hbin, bin)
	return err == nil
}

// 生成密码
func (this *Account) GeneratePassword(password string) string {
	bin := []byte(_accountSecret + password + this.Username)
	b, _ := bcrypt.GenerateFromPassword(bin, bcrypt.DefaultCost)
	this.Password = base64.StdEncoding.EncodeToString(b) // len = 80 = 64 hash 16 salt
	return this.Password
}

type AccountEmailVerify struct {
	Id        int64
	Code      string `qbs:"size:16,index,notnull"`
	AccountId int64  `qbs:"notnull"`
	Account   *Account
	Created   time.Time
}

func (this *AccountEmailVerify) GenerateCode() {
	this.Code = fmt.Sprintf("%016x", rand.Uint32())
}
