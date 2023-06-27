package parser

import (
	"HelloGo/crawler/engine"
	"log"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`class CC_DLL ([\S]+) ([: \S]+)
{
public:
	([\s\S^};]+)


};`)
var zhushiRe = regexp.MustCompile(`/\*\*[a-zA-Z0-9@*\s.:;(),-_/{}'=!&+\[\]"><\\]*\*/`)

// 正则匹配出公共成员函数
func ParseMemitemLocal(content []byte, url string) engine.ParseResult {
	// 替换掉注释
	//s := strings.Replace(string(content),string(96),"",-1) //96就是原生字符`
	//s = zhushiRe.ReplaceAllString(s, "")
	//if strings.Compare(url,"F:\\c++space\\cocos2d-x-3.17.2\\cocos\\2d\\CCNode.h") == 0{
	//	log.Printf("%s",url)
	//	log.Printf("%s",s) //调试
	//}

	submatch := memRe.FindAllSubmatch([]byte(content), -1)
	memCnt := 0
	result := engine.ParseResult{}

	//var members []string
	for _, m := range submatch {
		if strings.Compare(url, "F:\\c++space\\cocos2d-x-3.17.2\\cocos\\2d\\CCNode.h") == 0 {
			log.Printf("==%s", string(m[0]))
		}
		//	memitem := model.Memitem{}
		//	memitem.MemitemCenter = string(m[4])
		//	//构造、析构、重载函数冲突处理
		//	if memitem.MemitemCenter == classname || memitem.MemitemCenter == "~"+classname {
		//		continue
		//	}
		//	if strings.Contains(memitem.MemitemCenter, "~") {
		//		continue
		//	}
		//	idx := 0
		//	for _, m := range members {
		//		if strings.Contains(m, memitem.MemitemCenter) {
		//			idx++
		//		}
		//	}
		//	if idx > 0 { //同名函数加_序号
		//		memitem.MemitemCenter = fmt.Sprintf("%s_%d", memitem.MemitemCenter, idx)
		//	}
		//	members = append(members, memitem.MemitemCenter)
		//
		//	//url := string(m[2])
		//	//log.Printf("%s", right)
		//	//对返回值和参数再次解析
		//	paramSubmatch := returnRe.FindAllSubmatch(m[2], -1)
		//	var buffer bytes.Buffer
		//	for _, n := range paramSubmatch {
		//		//url := string(n[1])
		//		buffer.Write(n[1]) //返回值前缀如const等
		//		buffer.Write(n[4]) //返回值名
		//		buffer.Write(n[5]) //返回值后缀如* &等
		//	}
		//	memitem.MemitemLeft = strings.Replace(buffer.String(), "&amp;", "&", -1)
		//	memitem.MemitemLeft = strings.Replace(memitem.MemitemLeft, "&lt;", "<", -1)
		//	memitem.MemitemLeft = strings.Replace(memitem.MemitemLeft, "&gt;", ">", -1)
		//	buffer.Reset()
		//	items := strings.Split(string(m[5]), ",")
		//	for idx, item := range items {
		//		if idx != 0 {
		//			buffer.Write([]byte(","))
		//		}
		//		paramSubmatch = paramRe.FindAllSubmatch([]byte(item), -1)
		//		for _, n := range paramSubmatch {
		//			//url := string(n[1])
		//			buffer.Write(n[2])
		//			buffer.Write(n[3])
		//		}
		//	}
		//	memitem.MemitemRight = strings.Replace(buffer.String(), "&amp;", "&", -1)
		//	memitem.MemitemRight = strings.Replace(memitem.MemitemRight, "&lt;", "<", -1)
		//	memitem.MemitemRight = strings.Replace(memitem.MemitemRight, "&gt;", ">", -1)
		memCnt++
		//	//log.Printf("return: 		%s-->%s", classname, memitem.MemitemLeft)
		//	//log.Printf("%s: 	%s", classname,memitem.MemitemCenter)
		//	//log.Printf("params: 		%s", memitem.MemitemRight)
		//	memitem.ClassName = classname
		//	memitem.PackageName = packageName
		//	memitem.Namespace = namespace
		//	result.Items = append(result.Items, engine.Item{
		//		Url:     url,
		//		Payload: memitem,
		//	})
	}
	return result
}
