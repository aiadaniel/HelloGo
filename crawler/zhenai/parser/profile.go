package parser

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

/**
 * 解析每个人的信息
 */

func ParseProfile(contents []byte,name string,url string) engine.ParseResult {

	return parseFromHtml(contents,name,url)
}

//TODO: JSON方式获取
var profileRe = regexp.MustCompile(`<div class="m-btn [\w]+" data-v-bff6f798>([\S\p{Han}][^<]+)</div>`)
//性别目前需要专门获取
var genderRe = regexp.MustCompile(`"gender":([0-9]),"genderString":"([\p{Han}]+)"`)
var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseFromHtml(contents []byte,name string,url string) engine.ParseResult{
	all := profileRe.FindAllSubmatch(contents, -1)

	p := model.Profile{}
	p.Marriage = string(all[0][1])
	p.Age,_ = strconv.Atoi(strings.Trim(string(all[1][1]),"岁"))
	p.XingZuo = string(all[2][1])
	p.Name = name

	genderMatch := genderRe.FindAllSubmatch(contents, -1)
	for _,gender := range genderMatch {
		//fmt.Printf("get gender %s %s",gender[1],gender[2])
		p.Gender,_ = strconv.Atoi(string(gender[1]))
		p.GenderString = string(gender[2])
	}

	idMatch := idRe.FindAllSubmatch([]byte(url),-1)
	var id string
	for _,item := range idMatch{
		id = string(item[1])
	}

	for _,m := range all[3:] {
		//fmt.Printf("get %s\n",m[1])
		//网站的前3个目前有规律直接赋值
		item := string(m[1])
		if strings.ContainsAny(item,"cm") {
			p.Height,_ = strconv.Atoi(strings.Trim(item,"cm"))
			continue
		}
		if strings.ContainsAny(item,"kg") {
			p.Weight,_ = strconv.Atoi(strings.Trim(item,"kg"))
			continue
		}
		if strings.ContainsAny(item,"工作地") {
			//fmt.Println("工作地：" + item)
			continue
		}
		if strings.ContainsAny(item,"月收入") {
			p.Income = item
			continue
		}
		if strings.ContainsAny(item,"中专") ||
			strings.ContainsAny(item,"大专") ||
			strings.ContainsAny(item,"大学") ||
			strings.ContainsAny(item,"高中") ||
			strings.ContainsAny(item,"初中") ||
			strings.ContainsAny(item,"硕士") ||
			strings.ContainsAny(item,"博士") {
			p.Education = item
			continue
		}
		if strings.ContainsAny(item,"房") {
			p.House = item
			continue
		}
		if strings.ContainsAny(item,"车") {
			p.Car = item
			continue
		}
		//TODO 性别目前只在json段里看到

	}

	result := engine.ParseResult{
		Items:[]engine.Item{
			{
				Url: url,
				Type:"zhenai",
				Id:id,
				Payload:p,
			},
		},
	}
	return result
}