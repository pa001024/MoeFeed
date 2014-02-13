package models

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/pa001024/MoeWorker/util"
)

// 平台用户信息 多对一关联Account 实现单账号多用户
type PlatformUser struct {
	Id          int64
	AccountId   int64 `qbs:"unique,notnull"` // 关联账户
	Account     *Account
	DisplayName string `qbs:"size:64"`  // 显示名
	AvatarEmail string `qbs:"size:100"` // 头像邮箱
	Url         string `qbs:"size:100"` // 主页
	Status      int16  // 账户验证信息
	Created     time.Time
	Updated     time.Time
}

// enum User.Status
const (
	User     int16 = iota // 0: 未验证用户
	Team                  // 1: 组织
	SysAdmin              // 2: 鹳狸猿
)

// 获取头像地址
func (this *PlatformUser) AvatarUrl(size string) string {
	return fmt.Sprintf("https://secure.gravatar.com/avatar/%s?%s",
		util.Md5String(strings.ToLower(this.AvatarEmail)),
		(url.Values{
			"s": {size},
			"d": {"retro"}, // TODO: 自定义
		}).Encode())
}

// 返回字符串
func (this *PlatformUser) String() string {
	return fmt.Sprintf("PlatformUser(%s)", this.Account.Username)
}

// 返回登录时间
func (this *PlatformUser) Logined() string {
	return this.Updated.Format("2006年1月2日") // TODO: i18n
}

// 返回加入时间
func (this *PlatformUser) Joined() string {
	return this.Created.Format("2006年1月2日") // TODO: i18n
}

// 返回符号名称
func (this *PlatformUser) Name() string {
	return this.Account.Username
}

// 返回显示名称
func (this *PlatformUser) DispName() string {
	if this.DisplayName != "" {
		return this.DisplayName
	}
	return this.Account.Username
}
