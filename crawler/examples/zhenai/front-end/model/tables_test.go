package model

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/examples/zhenai/model"
	"html/template"
	"os"
	"strconv"
	"testing"
)

func TestTables(t *testing.T) {
	template := template.Must(template.ParseFiles("../template/tables.html"))

	file, err := os.Create("tables.test.html")
	if err != nil {
		panic(err)
	}

	datas := SearchResult{}
	datas.Hits = 2343
	for i := 0; i < 20; i++ {
		item := engine.Item{
			Url:  "http://album.zhenai.com/u/108906739",
			Type: "zhenai",
			Id:   "108906739",
			Payload: model.Profile{
				Name:         "珍珠",
				Gender:       1,
				GenderString: "女士",
				Age:          23,
				Height:       156,
				Weight:       65,
				Income:       strconv.Itoa(8990 + i),
				Marriage:     "已婚",
				Education:    "本科",
				Occupation:   "律师",
				Hukou:        "",
				XingZuo:      "白羊",
				House:        "已购房",
				Car:          "未购车",
			},
		}
		datas.Items = append(datas.Items, item)
	}
	template.Execute(file, datas) //os.Stdout
}
