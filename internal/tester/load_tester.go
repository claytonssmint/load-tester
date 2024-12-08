package tester

import (
	"net/http"
	"sync"
)

type LoadTestResults struct {
	TotalRequests int
	Status200     int
	OtherStatus   map[int]int
}

func RunLoadTest(url string, totalRequests, concurrency int) LoadTestResults {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := LoadTestResults{
		OtherStatus: make(map[int]int),
	}
	sem := make(chan struct{}, concurrency)
	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			defer mu.Unlock()
			results.TotalRequests++
			if resp.StatusCode == 200 {
				results.Status200++
			} else {
				results.OtherStatus[resp.StatusCode]++
			}

		}()
	}
	wg.Wait()
	return results

}
