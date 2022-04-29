package engine

import (
	"log"

	"github.com/Matafiya/webcrawler/crawler/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}

type Scheduler interface {
	Submit(Request)
	ReadyNotifier
	Run()
	WorkerChan() chan Request
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkCount; i++ {
		createWork(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {

			go func() { e.ItemChan <- item }()

		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

func createWork(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
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
