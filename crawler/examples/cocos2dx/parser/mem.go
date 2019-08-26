package parser

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/examples/cocos2dx/model"
	"log"
	"regexp"
)

//<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>
// 返回值类型一： <a class="el" href="../../d4/d5f/classcocos2d_1_1_scene.html">Scene</a> *&#160;
// 参数类型一： (<a class="el" href="../../d3/d82/classcocos2d_1_1_node.html">Node</a> *node)
var memRe = regexp.MustCompile(`<tr  name="cpp" class="memitem:[a-z0-9]+"><td class="memItemLeft" align="right" valign="top"><a class="anchor" id="[a-z0-9]+"></a>[\n\r]+([\S _]+)&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/classcocos2d_1_1_[a-zA-Z0-9]+.html)#[a-z0-9]+">([\S]+)</a> ([\S _]+)</td></tr>`)
var memRe2 = regexp.MustCompile(`<tr  name="cpp" class="memitem:[a-z0-9]+"><td class="memItemLeft" align="right" valign="top">([\S _]+)&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/classcocos2d_1_1_[a-zA-Z0-9]+.html)#[a-z0-9]+">([\S]+)</a> ([\S _]+)</td></tr>`)

var paramRe = regexp.MustCompile(`<a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/classcocos2d_1_1_[a-zA-Z0-9]+.html)">([\w]+)</a>([\s\*])`)

func ParseMemitem(bytes []byte, url string) engine.ParseResult {
	submatch := memRe.FindAllSubmatch(bytes, -1)
	memitem := model.Memitem{}
	memCnt := 0
	for _, m := range submatch {
		memitem.MemitemLeft = string(m[1])
		memitem.MemitemCenter = string(m[3])
		memitem.MemitemRight = string(m[4])
		//url := string(m[2])

		//对返回值和参数再次解析
		paramSubmatch := paramRe.FindAllSubmatch(m[1], -1)
		for _, n := range paramSubmatch {
			//url := string(n[1])
			memitem.MemitemLeft = string(n[2]) + string(n[3])
		}

		paramSubmatch = paramRe.FindAllSubmatch(m[4], -1)
		for _, n := range paramSubmatch {
			//url := string(n[1])
			memitem.MemitemRight = string(n[2]) + string(n[3])
		}

		memCnt++
		log.Printf("Public Member %02d %s %s%s", memCnt, memitem.MemitemLeft, memitem.MemitemCenter, memitem.MemitemRight)

	}
	submatch2 := memRe2.FindAllSubmatch(bytes, -1)
	for _, m := range submatch2 {
		memitem.MemitemLeft = string(m[1])
		memitem.MemitemCenter = string(m[3])
		memitem.MemitemRight = string(m[4])
		//url := string(m[2])
		memCnt++
		log.Printf("Public Member %02d %s %s%s", memCnt, memitem.MemitemLeft, memitem.MemitemCenter, memitem.MemitemRight)

	}
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Payload: memitem,
			},
		},
	}

	return result
}
