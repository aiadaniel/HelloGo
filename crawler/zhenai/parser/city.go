package parser

import (
	"HelloGo/crawler/engine"
	"regexp"
)

/**
 * 解析城市首页用户列表
 */

var (
	pRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	//下一页 获取更多城市
	cityRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(bytes []byte,url string) engine.ParseResult {
	submatchs := pRe.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}
	for _, m := range submatchs {
		url := string(m[1])
		name := string(m[2])
		//fmt.Printf("user:%s url:%s\n", name, m[1])
		//result.Items = append(result.Items, "user "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte,url string) engine.ParseResult {
				return ParseProfile(c, name,url)
			},
		})
	}

	nextcitys := cityRe.FindAllSubmatch(bytes, -1)
	for _,m := range nextcitys {
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCity,
		})
	}

	//fmt.Printf("ALL match size %d\n", len(submatchs))
	return result
}
