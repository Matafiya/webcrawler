package parser

import (
	"regexp"

	"github.com/Matafiya/webcrawler/crawler/engine"
)

const cityListRe = `<a href="(http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {

		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: engine.NilParser,
		})

	}

	return result
}
