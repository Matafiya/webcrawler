package engine

import (
	"log"

	"github.com/Matafiya/webcrawler/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error fetureing url %s: %v", r.URL, err)
			continue
		}
		result := r.ParserFunc(body)
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
		}

	}
}
