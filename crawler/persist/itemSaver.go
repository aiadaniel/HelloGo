package persist

import (
	"HelloGo/crawler/engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item,error) {
	out := make(chan engine.Item)
	//当前测试无集群
	client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL("http://192.168.128.129:9200"))
	if err != nil {
		return nil,err
	}
	go func() {
		itemCnt := 0
		for {
			item := <-out
			log.Printf("ItemSaver got:%d, %v", itemCnt, item)
			itemCnt++

			err := saveItem(client,index,item)
			if err != nil {
				log.Printf("err save %v %d",item,err)
			}

		}
	}()
	return out,nil
}

//elastic 几个核心概念
//index 相当于数据库名
//type 表名
//id
func saveItem(client *elastic.Client, index string,item engine.Item) error {

	if item.Type == "" {
		return errors.New("must set type")
	}
	indexService := client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf(response)
	return nil
}
