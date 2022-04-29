package main

import (
	"github.com/Matafiya/webcrawler/crawler/engine"
	"github.com/Matafiya/webcrawler/crawler/persist"
	"github.com/Matafiya/webcrawler/crawler/scheduler"
	"github.com/Matafiya/webcrawler/crawler/zhenai/parser"
)

func main() {
	/*
		engine.SimpleEngine{}.Run(engine.Request{
			URL:        "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		})
	*/
	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
		ItemChan:  persist.ItemSaver()}
	e.Run(engine.Request{
		URL:        "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
