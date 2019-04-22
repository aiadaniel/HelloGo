package parser

import (
	"HelloGo/crawler/engine"
	"regexp"
)

/**
 * 解析城市列表
 */

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(bytes []byte,url string) engine.ParseResult {
	//re := regexp.MustCompile(cityListRe)
	submatchs := cityListRe.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}
	for _, m := range submatchs {
		//fmt.Printf("city:%s url:%s\n", m[2], m[1])
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		//测试，先只拿一个城市，否则太久
		//return result
	}

	//fmt.Printf("ALL match size %d\n", len(submatchs))
	return result
}
