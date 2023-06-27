package main

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/engine/scheduler"
	"HelloGo/crawler/parsers/cocos2dx/parser"
	"HelloGo/crawler/persist/cocos2dx"
)

//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/db/d59/classcocos2d_1_1_grid3_d.html"
//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/d7/df3/classcocos2d_1_1_director.html"
//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/d4/d5f/classcocos2d_1_1_scene.html"

//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/d0/dc7/group____2d.html"
//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/dd/df2/group__base.html"
const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/da/d35/group____3d.html"

//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/dd/d0d/group__actions.html"
//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/d0/d2f/group__physics__2d.html"
//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/dd/df1/group__ui.html"
//const urlcocos string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x/dd/d00/group__renderer.html"

// NOTE: 这种api爬虫的方式，有个问题就是api文档的更新较慢，有时跟代码对应不上，所以换本地文件爬虫的方式。
func main_1() {
	//en := engine.SimpleEngine{}
	//太快的请求触发爬虫网站的防御机制，导致403
	itemChan, err := cocos2dx.ItemSaverCocos("")
	if err != nil {
		panic(err)
	}
	en := engine.ConcurrentEngine{
		WorkerCnt: 16,
		Scheduler: &scheduler.QueueScheduler{},
		ItemChan:  itemChan,
	}
	en.Run(false, engine.Request{
		Url:        urlcocos,
		ParserFunc: parser.ParseCocosPackage,
	})
}
