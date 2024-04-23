package main

import (
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/superjcd/gocrawler/request"
	"github.com/superjcd/gocrawler/scheduler"
	"github.com/superjcd/gocrawler/scheduler/nsq"
)

func main() {
	s := nsq.NewNsqScheduler("zyte", "default", "127.0.0.1:4150", "127.0.0.1:4161")
	pages := []int{}
	for i := 1; i < 10; i++ {
		pages = append(pages, i)
	}
	uid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	log.Printf("taskId: %s", uid.String())

	for _, pg := range pages {
		data := make(map[string]string, 0)
		data["taskId"] = uid.String()
		url := fmt.Sprintf("https://www.zyte.com/blog/page/%d", pg)
		fmt.Println(url)
		req := request.Request{
			URL:    url,
			Method: "GET",
			Data:   data,
		}
		s.Push(scheduler.TYP_PUSH_SCHEDULER, &req)

	}
}
