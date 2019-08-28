package engine

import (
	"HelloGo/crawler/fetcher"
	"fmt"
	"time"
)

var rateLimiter = time.Tick(50 * time.Millisecond)

func Worker(r Request) (ParseResult, error) {
	<-rateLimiter
	//fmt.Printf("Fetching %s\n",r.Url)
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		fmt.Printf("err fetching:%s %v\n", r.Url, err)
		return ParseResult{}, nil
	}
	//log.Printf("%d %s", len(contents),string(contents[5500:6000]))

	return r.ParserFunc(contents, r.Url), nil
}
