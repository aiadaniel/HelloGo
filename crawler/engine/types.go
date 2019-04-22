package engine

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type ParserFunc func(contents []byte,url string) ParseResult
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type Engine interface {
	Run(seeds ...Request)
}

func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
