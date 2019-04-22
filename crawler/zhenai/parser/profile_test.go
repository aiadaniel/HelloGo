package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseFromHtml(t *testing.T)  {
	file, err := ioutil.ReadFile("profile_test_data2.html")
	if err != nil {
		panic(err)
	}
	parseFromHtml(file,"haha")
}
