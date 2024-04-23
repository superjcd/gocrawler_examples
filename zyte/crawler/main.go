package main

import (
	default_builder "github.com/superjcd/gocrawler/builder/default"
	"github.com/superjcd/gocrawler_examples/zyte/parser"
)

func main() {
	config := default_builder.DefaultWorkerBuilderConfig{}
	worker := config.Name("zyte").MaxRunTime(300).Workers(10).LimitRate(10).Build(parser.NewZyteParser())
	worker.Run()
}
