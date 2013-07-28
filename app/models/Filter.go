package models

import (
	"encoding/json"
	"sort"
)

// 过滤器
type Filter JobData

const (
	FilterText   = iota // [检测系]文本过滤
	FilterRegexp        // [检测系]正则过滤
)
const (
	FilterMediawikiApi = 100 + iota // [抓取系]获取WikiText
)
const (
	FilterZhconv          = 200 + iota // [转换系]简繁转换
	FilterMachineTranlate              // [转换系]机器翻译
	FilterHumanTranlate                // [转换系]人工翻译
	FilterWikitext                     // [转换系]WIKI文本过滤
	FilterHtmltext                     // [转换系]HTML文本过滤
	FilterExecText                     // [转换系]文本替换
	FilterExecRegexp                   // [转换系]正则替换
	FilterExecJavascript               // [转换系]执行js代码
)
const (
	FilterTextRender  = 300 + iota // [渲染系]格式化文本
	FilterImageRender              // [渲染系]按模板合成图片与文字
)

func (this *Filter) ViewData() (rst []KeyPair) {
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
