package parser

import (
	"regexp"

	"github.com/Matafiya/webcrawler/crawler/engine"
)

const cityRe = `<a href="(http://127.0.0.1:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				URL: string(m[1]),
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseProfile(contents, name)
				},
			})
	}

	return result
}
