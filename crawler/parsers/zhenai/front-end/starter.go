package main

import (
	"HelloGo/crawler/parsers/zhenai/front-end/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/front-end/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler("crawler/front-end/view/template/tables.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
