package main

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/persist"
	"HelloGo/crawler/scheduler"
	"HelloGo/crawler/zhenai/parser"
)

const url string = "http://www.zhenai.com/zhenghun"

func main() {
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
	en.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
