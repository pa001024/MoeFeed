package models

import (
	"encoding/json"
	"sort"
)

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

func (this *Source) ViewData() (rst []KeyPair) {
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
