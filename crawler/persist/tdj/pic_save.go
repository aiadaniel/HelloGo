package tdj_persist

import (
	"HelloGo/crawler/engine"
	"os"
	"path/filepath"
	"strings"
)

func PicSave() (chan engine.Item, error) {
	out := make(chan engine.Item)

	go func() {
		itemCnt := 0
		for {
			item := <-out
			// log.Printf("ItemSaver got:%d, %v", itemCnt, item)
			itemCnt++

			paths := strings.Split(item.Url, "/")
			imgPath := paths[len(paths)-1]
			imgPath = filepath.Join(".", imgPath)
			file, _ := os.Create(imgPath)
			content := item.Payload.([]byte)
			file.Write(content)

			file.Close()

		}
	}()
	return out, nil
}
