package parser

import (
	"HelloGo/crawler/engine"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 遍历文件夹
var dirs = make([]string, 128)

func ParseCocosPackageLocal(content []byte, url1 string) engine.ParseResult {

	err := filepath.Walk(url1, myWalkFunc)
	if err != nil {
		log.Printf("PackageLocal err: %v", err)
	}

	result := engine.ParseResult{}
	for _, m := range dirs {
		url := m
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte, url2 string) engine.ParseResult {
				return ParseMemitemLocal(c, url)
			},
		})
		//log.Printf("=====%s %s %s",packageName,namespaceName,moduleName)
	}
	return result
}

const seperator = string(os.PathSeparator)

func myWalkFunc(path string, info os.FileInfo, err error) error {

	if strings.Contains(path, "scripting"+seperator) ||
		strings.Contains(path, "storage"+seperator) ||
		strings.Contains(path, "UIEditBox"+seperator) ||
		strings.Contains(path, "docs"+seperator) ||
		strings.Contains(path, "external"+seperator) ||
		strings.Contains(path, "licenses"+seperator) ||
		strings.Contains(path, "templates"+seperator) ||
		strings.Contains(path, "tests"+seperator) ||
		strings.Contains(path, "tools"+seperator) ||
		strings.Contains(path, "web"+seperator) ||
		strings.Contains(path, "build"+seperator) ||
		strings.Contains(path, "cmake"+seperator) ||
		strings.Contains(path, "audio"+seperator) ||
		strings.Contains(path, "allocator"+seperator) ||
		strings.Contains(path, "cocos2d.h") ||
		strings.Contains(path, "deprecated"+seperator) ||
		strings.Contains(path, "editor-support"+seperator) ||
		strings.Contains(path, "-android") ||
		strings.Contains(path, "-apple") ||
		strings.Contains(path, "-winrt") ||
		strings.Contains(path, "-ios") ||
		strings.Contains(path, "platform"+seperator) ||
		strings.Contains(path, ".py") ||
		strings.Contains(path, ".txt") ||
		strings.Contains(path, ".xml") ||
		strings.Contains(path, ".cpp") ||
		strings.Contains(path, ".frag") ||
		strings.Contains(path, ".md") {

	} else {
		if strings.HasSuffix(path, ".h") {
			//log.Printf("now: %s", path)
			dirs = append(dirs, path)
		}
	}
	return nil
}
