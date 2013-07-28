package models

import (
	"encoding/json"
	"sort"
)

// 目标
type Target JobData

const (
	TargetSinaWeibo    = iota // [时间线系]新浪微博
	TargetTencentWeibo        // [时间线系]腾讯微博
	Target163Weibo            // [时间线系]网易微博
	TargetSohoWeibo           // [时间线系]搜狐微博
	TargetFacebook            // [时间线系]脸书
	TargetTwitter             // [时间线系]推特
	TargetRenren              // [时间线系]人人
	TargetKaixin              // [时间线系]开心网
	TargetDouban              // [时间线系]豆瓣
	TargetWeimoe              // [时间线系]微萌
)
const (
	TargetBlog       = 100 + iota // [博客系]博客
	TargetWordpress               // [博客系]WP博客
	TargetEmlog                   // [博客系]EL博客
	TargetBaiduSpace              // [博客系]百度空间
	TargetBaiduTieba              // [博客系]贴吧
	Target163Blog                 // [博客系]网易轻博客
	TargetQzone                   // [博客系]QQ空间
)
const (
	TargetBBS     = 200 + iota // [讨论版系]论坛
	TargetDiscuz               // [讨论版系]DZ论坛
	TargetPHPWind              // [讨论版系]PW论坛
	TargetXenForo              // [讨论版系]XF论坛
	TargetHAcfun               // [讨论版系]AC匿名区
)
const (
	TargetQQSingle = 300 + iota // [IM系]QQ私聊
	TargetQQGroup               // [IM系]QQ群
	TargetIRC                   // [IM系]IRC
	TargetXMPP                  // [IM系]XMPP
)
const (
	TargetSendMail    = 400 + iota // [服务系]邮件
	TargetHttpRequest              // [服务系]HTTP
)

func (this *Target) ViewData() (rst []KeyPair) {
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
