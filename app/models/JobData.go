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
	Name      string `qbs:"size:32,notnull"`
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

// 来源
type Source JobData

const (
	SourceRSS  = iota // [主动系]RSS订阅
	SourceAtom        // [主动系]Atom订阅
	SourceHtml        // [主动系]Html订阅
)
const (
	SourceSinaWeibo    = 100 + iota // [时间线系]新浪微博
	SourceTencentWeibo              // [时间线系]疼迅微博
	Source163Weibo                  // [时间线系]网易微博
	SourceSohoWeibo                 // [时间线系]搜狐微博
	SourceFacebook                  // [时间线系]脸书
	SourceTwitter                   // [时间线系]推特
	SourceRenren                    // [时间线系]人人
	SourceKaixin                    // [时间线系]开心网
	SourceDouban                    // [时间线系]豆瓣
	SourceWeimoe                    // [时间线系]微萌
)
const (
	SourceListenHttp = 200 + iota // [被动系]Http订阅
	SourceGithubHook
)

// 过滤器
type Filter JobData

const (
	FilterText   = iota // [检测系]文本过滤
	FilterRegexp        // [检测系]文本
)
const (
	FilterWikiApi = 100 + iota // [抓取系]获取WikiText
)
const (
	FilterMachineTranlate = 200 + iota // [转换系]机器翻译
	FilterHumanTranlate                // [转换系]人工翻译
	FilterZhconv                       // [转换系]简繁转换
	FilterWikitext                     // [转换系]WIKI文本过滤
	FilterHtmltext                     // [转换系]HTML文本过滤
	FilterExecText                     // [转换系]文本替换
	FilterExecRegexp                   // [转换系]正则替换
	FilterExecJavascript               // [转换系]执行js代码
)
const (
	FilterImageRender = 300 + iota // [渲染系]按模板合成图片与文字
	FilterImageMerger              // [渲染系]按程序合并图片
)

// 目标
type Target JobData

const (
	TargetSinaWeibo    = iota // [时间线系]新浪微博
	TargetTencentWeibo        // [时间线系]疼迅微博
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
