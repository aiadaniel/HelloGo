package main

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/engine/scheduler"
	"HelloGo/crawler/parsers/zhenai/parser"
	"HelloGo/crawler/persist"
)

const urlzhenai string = "http://www.zhenai.com/zhenghun"

func main_zhenai() {
	//en := engine.SimpleEngine{}
	//太快的请求触发爬虫网站的防御机制，导致403
	itemChan, err := persist.ItemSaver("data_profile")
	if err != nil {
		panic(err)
	}
	en := engine.ConcurrentEngine{
		WorkerCnt: 50,
		Scheduler: &scheduler.QueueScheduler{},
		ItemChan:  itemChan,
	}
	en.Run(false, false, engine.Request{
		Url:        urlzhenai,
		ParserFunc: parser.ParseCityList,
	})
}
