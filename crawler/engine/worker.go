package engine

import (
	"HelloGo/crawler/engine/fetcher"
	"fmt"
	"time"
)

var rateLimiter = time.Tick(50 * time.Millisecond)

// isLocal表示是否本地文件爬取
// download表示是否下载文件而不是读取文件内容
func Worker(r Request, isLocal bool, download bool) (ParseResult, error) {
	<-rateLimiter
	//fmt.Printf("Fetching %s\n",r.Url)
	var contents []byte
	var err error
	if isLocal {
		contents, err = fetcher.FetchLocal(r.Url)
	} else {
		contents, err = fetcher.Fetch(r.Url, download)
	}
	if err != nil {
		fmt.Printf("err fetching:%s %v\n", r.Url, err)
		return ParseResult{}, nil
	}
	//log.Printf("%d %s", len(contents),string(contents[5500:6000]))

	return r.ParserFunc(contents, r.Url), nil
}
