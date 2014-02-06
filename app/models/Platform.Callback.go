package models

import (
	r "github.com/robfig/revel"
	"time"
)

// 路线
type Callback struct {
	Id        int64
	Name      string `qbs:"size:32,notnull"`
	Url       string `qbs:"size:32,index,notnull"`
	Type      int16  `qbs:"notnull"`
	ProjectId int64  `qbs:"notnull"`
	Project   *Project
	Created   time.Time
	Updated   time.Time
}

// 传递式验证
func (this *Callback) Validate(v *r.Validation, password string) {
	v.Check(this.Name, r.Required{}, r.MinSize{2}, r.MaxSize{32})
	v.Check(this.Url, r.Required{}, r.MinSize{2}, r.MaxSize{32})
}

const (
	CallbackOAuthSinaWeibo    = iota // [OAuth系]新浪微博OAuth2.0
	CallbackOAuthTencentWeibo        // [OAuth系]腾讯微博OAuth2.0
	CallbackOAuth163Weibo            // [OAuth系]网易微博OAuth2.0
	CallbackOAuthSohoWeibo           // [OAuth系]搜狐微博OAuth2.0
	CallbackOAuthFacebook            // [OAuth系]脸书
	CallbackOAuthTwitter             // [OAuth系]推特
	CallbackOAuthRenren              // [OAuth系]人人
	CallbackOAuthKaixin              // [OAuth系]开心网
	CallbackOAuthDouban              // [OAuth系]豆瓣
	CallbackOAuthWeimoe              // [OAuth系]微萌
)

const (
	CallbackHttp      = 100 + iota // [Ping系]HTTP
	CallbackWebSocket              // [Ping系]WebSocket
)

const (
	CallbackPullReply  = 200 + iota // [响应系]自动回复
	CallbackPullSearch              // [响应系]Feed搜索
	CallbackPullStatus              // [响应系]状态
)
