package engine

import (
	"log"

	"github.com/Matafiya/webcrawler/crawler/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerReady(chan Request)
	Run()
	WorkerChan() chan Request
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)

	for i := 0; i < e.WorkCount; i++ {
		createWork(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item : %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

func createWork(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetureing url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
