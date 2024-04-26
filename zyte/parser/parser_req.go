package parser

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofrs/uuid"
	"github.com/superjcd/gocrawler/parser"
	"github.com/superjcd/gocrawler/request"
)

type zyteReqParser struct{}

func NewZyteReqParser() *zyteReqParser {
	return &zyteReqParser{}
}

func (p *zyteReqParser) Parse(ctx context.Context, r *http.Response) (*parser.ParseResult, error) {
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, err
	}
	result := &parser.ParseResult{}
	resultItems := make([]parser.ParseItem, 0)
	requests := []*request.Request{}
	ctxValue := ctx.Value(request.RequestDataCtxKey{})
	requestData := ctxValue.(map[string]string)
	page := requestData["page"]

	if page == "1" {
		uid, _ := uuid.NewV4()
		for pg := 2; pg <= 5; pg++ {
			data := make(map[string]string, 0)
			data["taskId"] = uid.String()
			data["page"] = strconv.Itoa(pg)
			url := fmt.Sprintf("https://www.zyte.com/blog/page/%d", pg)
			req := request.Request{
				URL:    url,
				Method: "GET",
				Data:   data,
			}
			requests = append(requests, &req)
		}
	}

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
	result.Requests = requests

	return result, nil
}
