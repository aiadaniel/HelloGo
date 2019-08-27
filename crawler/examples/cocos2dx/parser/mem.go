package parser

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/examples/cocos2dx/model"
	"bytes"
	"log"
	"regexp"
	"strings"
)

var memRe = regexp.MustCompile(
					`<tr  name="cpp" class="memitem:[a-z0-9]+"><td class="memItemLeft" align="right" valign="top">(<a class="anchor" id="[a-z0-9]+"></a>[\n\r]+)*([\S _]+)&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/classcocos2d_1_1_[a-zA-Z0-9]+.html)#[a-z0-9]+">([\S]+)</a>([\S _]+)</td></tr>`)
var returnRe = regexp.MustCompile( //返回值
					`(a-zA-Z0-9 :&;)*(<a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/[a-zA-Z0-9_]+.html[#a-z0-9]*)">([a-zA-Z0-9 _]+)</a>)*([a-zA-Z0-9 :&;*]*)`) //&#160;
var paramRe = regexp.MustCompile( //参数
	`(<a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/[a-zA-Z0-9_]+.html[#a-z0-9]*)">([a-zA-Z0-9 _]+)</a>)*([a-zA-Z0-9 ():&;*=]*)`) //&#160;
func ParseMemitem(content []byte, url string) engine.ParseResult {
	submatch := memRe.FindAllSubmatch(content, -1)
	memitem := model.Memitem{}
	memCnt := 0
	for _, m := range submatch {
		left := string(m[2])
		right := string(m[5])
		memitem.MemitemCenter = string(m[4])
		//url := string(m[2])
		//log.Printf("%s",right)
		//对返回值和参数再次解析
		paramSubmatch := returnRe.FindAllSubmatch([]byte(left), -1)
		var buffer bytes.Buffer
		for _, n := range paramSubmatch {
			//url := string(n[1])
			buffer.Write(n[1]) //返回值前缀如const等
			buffer.Write(n[4]) //返回值名
			buffer.Write(n[5]) //返回值后缀如* &等
		}
		memitem.MemitemLeft = strings.Replace(buffer.String(), "&amp;", "&", -1)
		buffer.Reset()
		items := strings.Split(right, ",")
		for idx, item := range items {
			if idx != 0 {
				buffer.Write([]byte(","))
			}
			paramSubmatch = paramRe.FindAllSubmatch([]byte(item), -1)
			for _, n := range paramSubmatch {
				//url := string(n[1])
				buffer.Write(n[3])
				buffer.Write(n[4])
				//buffer.Write(n[5])
			}
		}
		memitem.MemitemRight = strings.Replace(buffer.String(), "&amp;", "&", -1)
		memCnt++
		log.Printf("==>%02d %s %s %s",
			memCnt, memitem.MemitemLeft, memitem.MemitemCenter, memitem.MemitemRight)
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
