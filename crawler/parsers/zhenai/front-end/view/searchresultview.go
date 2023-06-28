package view

import (
	"HelloGo/crawler/parsers/zhenai/front-end/model"
	"html/template"
	"net/http"
)

type SearchResultView struct {
	t *template.Template
}

func CreateSearchResultView(t string) SearchResultView {
	template, err := template.ParseFiles(t)
	if err != nil {
		panic(err)
	}

	return SearchResultView{
		t: template,
	}
}

func (v *SearchResultView) Render(w http.ResponseWriter, data model.SearchResult) error {
	err := v.t.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
