package parser

import (
	"testing"

	"github.com/Matafiya/webcrawler/crawler/fetcher"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d", resultSize, len(result.Requests))
	}

	var expectedUrls = []string{
		"http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}

	var expectedCites = []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s;but "+"was %s", i, url, result.Requests[i].URL)
		}
	}

	for i, city := range expectedCites {
		if result.Items[i] != city {
			t.Errorf("expected city #%d: %s;but "+"was %s", i, city, result.Items[i])
		}
	}

}
