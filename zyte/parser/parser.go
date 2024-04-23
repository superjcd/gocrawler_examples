package parser

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/superjcd/gocrawler/parser"
)

type zyteParser struct{}

func NewZyteParser() *zyteParser {
	return &zyteParser{}
}

func (p *zyteParser) Parse(ctx context.Context, r *http.Response) (*parser.ParseResult, error) {
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, err
	}
	result := &parser.ParseResult{}
	resultItems := make([]parser.ParseItem, 0)

	doc.Find("div.CardResource_card__BhCok").Each(
		func(i int, s *goquery.Selection) {
			item := parser.ParseItem{}
			item["title"] = s.Find("div.free-text").Text()
			item["author"] = s.Find("div:nth-child(3) > div:nth-child(1) > span:nth-child(2)").Text()
			item["read_time"] = s.Find("div:nth-child(3) > div:nth-child(2) > span:nth-child(2)").Text()
			item["post_time"] = s.Find("div:nth-child(4) > div:nth-child(1) > span:nth-child(2)").Text()
			resultItems = append(resultItems, item)
		},
	)
	result.Items = resultItems

	return result, nil
}
