package persist

import (
	"HelloGo/crawler/engine"
)

func ItemSaverPrint(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		itemCnt := 0
		for {
			_ = <-out
			//item := <-out
			//log.Printf("ItemSaver got:%d, %v", itemCnt, item)
			itemCnt++

		}
	}()
	return out, nil
}
