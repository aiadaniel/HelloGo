package parser

import (
	"HelloGo/crawler/engine"
	"log"
	"regexp"
	"strings"
)

const baseurl string = "https://docs.cocos2d-x.org/api-ref/cplusplus/v3x"

var packageNameRe = regexp.MustCompile(`<div class="title">([a-zA-Z0-9 ]+)</div>`)

var namespaceRe = regexp.MustCompile(
	`href="../../[a-z0-9]+/[a-z0-9]+/namespacecocos2d[a-zA-Z0-9_]+.html">([a-zA-Z0-9:]+)</a></td></tr>`)

var classesRe = regexp.MustCompile(
	`<tr class="memitem:"><td class="memItemLeft" align="right" valign="top">([a-zA-Z]+) &#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/[a-zA-Z0-9_]+.html)">([a-zA-Z0-9]+)</a></td></tr>`)

// 模块主页，内部包含类列表
func ParseCocosPackage(content []byte, url1 string) engine.ParseResult {
	submatch := packageNameRe.FindAllSubmatch(content, -1)
	var packageName, namespaceName string
	for _, m := range submatch {
		packageName = string(m[1])
	}

	allSubmatch := namespaceRe.FindAllSubmatch(content, -1)
	for _, m := range allSubmatch {
		namespaceName = string(m[1])
	}
	log.Printf("=====package:%s namespace:%s", packageName, namespaceName)

	result := engine.ParseResult{}
	classes := classesRe.FindAllSubmatch(content, -1)
	for _, m := range classes {
		classorstruct := string(m[1])
		if strings.Compare(classorstruct, "struct") == 0 ||
			strings.Compare(classorstruct, "Clonable") == 0 {
			continue
		}

		var className = string(m[3])
		//log.Printf("==class: %s %s", className, string(m[2]))
		var url = baseurl + string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte, url2 string) engine.ParseResult {
				return ParseMemitem(c, url, packageName, className, namespaceName)
			},
		})
		//log.Printf("=====%s %s %s",packageName,namespaceName,moduleName)
	}

	return result
}
