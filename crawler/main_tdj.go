package main

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/engine/scheduler"
	"HelloGo/crawler/parsers/tdj"
	tdj_persist "HelloGo/crawler/persist/tdj"
	"fmt"
)

// 目前实现爬取天地劫官方网站的角色介绍图片，
// 网址是规律的，即http://123.206.184.101/media/pictures/cn/tdj/img/new/hero/h1_peo.png 这个后缀按数字递增即可，所以不需要其他解析了，直接调度下载
const tdjpic_pre string = "http://123.206.184.101/media/pictures/cn/tdj/img/new/hero/h"
const tdjpic_suf string = "_peo.png"

const MaxPic int = 78 //官网目前测试有这么多张

func main_tdj() {
	itemChan, err := tdj_persist.PicSave()
	if err != nil {
		panic(err)
	}
	en := engine.ConcurrentEngine{
		WorkerCnt: 50,
		Scheduler: &scheduler.QueueScheduler{},
		ItemChan:  itemChan,
	}
	var reqs []engine.Request
	var i int
	for i = 0; i < MaxPic; i++ {
		reqs = append(reqs, engine.Request{
			Url:        fmt.Sprintf("%s%d%s", tdjpic_pre, i, tdjpic_suf),
			ParserFunc: tdj.ParsePic,
		})
	}
	en.Run(false, reqs...)
}
