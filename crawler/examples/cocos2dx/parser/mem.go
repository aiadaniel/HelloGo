package parser

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/examples/cocos2dx/model"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var memRe = regexp.MustCompile(
					`<tr  name="cpp" class="memitem:[a-z0-9]+"><td class="memItemLeft" align="right" valign="top">(<a class="anchor" id="[a-z0-9]+"></a>[\s]+)*([\S _]+)&#160;</td><td class="memItemRight" valign="bottom"><a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/classcocos2d_1_1_[a-zA-Z0-9]+.html)#[a-z0-9]+">([\S]+)</a>([\S _]+)</td></tr>`)
var returnRe = regexp.MustCompile( //返回值
					`(a-zA-Z0-9 :&;)*(<a class="el" href="../..(/[a-z0-9]+/[a-z0-9]+/[a-zA-Z0-9_]+.html[#a-z0-9]*)">([a-zA-Z0-9 _]+)</a>)*([a-zA-Z0-9 :&;*]*)`) //&#160;
var paramRe = regexp.MustCompile( //参数
	`(<a class="el" href="../../[a-z0-9]+/[a-z0-9]+/[a-zA-Z0-9_]+.html[#a-z0-9]*">([a-zA-Z0-9 _]+)</a>)*([a-zA-Z0-9 :&;*=]*)`) //&#160;
// 具体类页面
func ParseMemitem(content []byte, url string, packageName string, classname string, namespace string) engine.ParseResult {
	submatch := memRe.FindAllSubmatch(content, -1)
	memCnt := 0
	result := engine.ParseResult{}
	//if len(submatch) == 0 {
	//	panic("may be something err on url")
	//}
	var members []string
	for _, m := range submatch {
		memitem := model.Memitem{}
		memitem.MemitemCenter = string(m[4])
		//构造、析构、重载函数冲突处理
		if memitem.MemitemCenter == classname || memitem.MemitemCenter == "~"+classname {
			continue
		}
		if strings.Contains(memitem.MemitemCenter, "~") {
			continue
		}
		idx := 0
		for _, m := range members {
			if strings.Contains(m, memitem.MemitemCenter) {
				idx++
			}
		}
		if idx > 0 { //同名函数加_序号
			memitem.MemitemCenter = fmt.Sprintf("%s_%d", memitem.MemitemCenter, idx)
		}
		members = append(members, memitem.MemitemCenter)

		//url := string(m[2])
		//log.Printf("%s", right)
		//对返回值和参数再次解析
		paramSubmatch := returnRe.FindAllSubmatch(m[2], -1)
		var buffer bytes.Buffer
		for _, n := range paramSubmatch {
			//url := string(n[1])
			buffer.Write(n[1]) //返回值前缀如const等
			buffer.Write(n[4]) //返回值名
			buffer.Write(n[5]) //返回值后缀如* &等
		}
		memitem.MemitemLeft = strings.Replace(buffer.String(), "&amp;", "&", -1)
		memitem.MemitemLeft = strings.Replace(memitem.MemitemLeft, "&lt;", "<", -1)
		memitem.MemitemLeft = strings.Replace(memitem.MemitemLeft, "&gt;", ">", -1)
		buffer.Reset()
		items := strings.Split(string(m[5]), ",")
		for idx, item := range items {
			if idx != 0 {
				buffer.Write([]byte(","))
			}
			paramSubmatch = paramRe.FindAllSubmatch([]byte(item), -1)
			for _, n := range paramSubmatch {
				//url := string(n[1])
				buffer.Write(n[2])
				buffer.Write(n[3])
				//TODO (<a class="el" href="../../dd/df2/group__base.html#ga7885f47644a0388f981f416fa20389b2">EventKeyboard::KeyCode</a> keyCode 匹配仍然有问题
			}
		}
		memitem.MemitemRight = strings.Replace(buffer.String(), "&amp;", "&", -1)
		memitem.MemitemRight = strings.Replace(memitem.MemitemRight, "&lt;", "<", -1)
		memitem.MemitemRight = strings.Replace(memitem.MemitemRight, "&gt;", ">", -1)
		memCnt++
		//log.Printf("return: 		%s-->%s", classname, memitem.MemitemLeft)
		//log.Printf("%s: 	%s", classname,memitem.MemitemCenter)
		//log.Printf("params: 		%s", memitem.MemitemRight)
		memitem.ClassName = classname
		memitem.PackageName = packageName
		memitem.Namespace = namespace
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Payload: memitem,
		})
	}
	return result
}
