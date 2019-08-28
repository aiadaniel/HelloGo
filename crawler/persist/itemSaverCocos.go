package persist

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/examples/cocos2dx/model"
	"log"
)

/**
 * 爬虫收到的解析包使用模板引擎或者字符串直接处理的方式，生成头文件和.cpp文件
 */
func ItemSaverCocos(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		itemCnt := 0
		for {
			//_ = <-out
			item := <-out
			var memFunc = item.Payload.(model.Memitem)
			log.Printf("ItemSaver got:%d, %s", itemCnt, memFunc)
			itemCnt++
		}
	}()
	return out, nil
}
