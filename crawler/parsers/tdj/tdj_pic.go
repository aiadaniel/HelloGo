package tdj

import "HelloGo/crawler/engine"

func ParsePic(contents []byte, url string) engine.ParseResult {
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "tdj",
				Id:      "id",
				Payload: contents,
			},
		},
	}
	return result
}
