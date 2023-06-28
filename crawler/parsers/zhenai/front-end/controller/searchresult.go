package controller

import (
	"HelloGo/crawler/engine"
	"HelloGo/crawler/parsers/zhenai/front-end/model"
	"HelloGo/crawler/parsers/zhenai/front-end/view"
	"context"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.128.129:9200"))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	trimSpace := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	result, err := h.getSearchResult(trimSpace, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
func (h SearchResultHandler) getSearchResult(q string, from int) (page model.SearchResult, err error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.Search("data_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil

}

//对于如输入参数Age<30 自动转成Payload:Age<30
func rewriteQueryString(q string) string {
	compile := regexp.MustCompile(`([A-Z][a-z]*):`)
	return compile.ReplaceAllString(q, "Payload.$1:")
}
