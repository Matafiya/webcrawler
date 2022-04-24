package main

import (
	"github.com/Matafiya/webcrawler/crawler/engine"
	"github.com/Matafiya/webcrawler/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
