package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m   map[string]bool
	mux sync.Mutex
}

func (s *SafeMap) Add(str string) {
	s.mux.Lock()
	s.m[str] = true
	s.mux.Unlock()
}
func (s *SafeMap) Contains(str string) bool {
	s.mux.Lock()
	_, res := s.m[str]
	s.mux.Unlock()
	return res
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	msg := make(chan string)
	s := SafeMap{m: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	go _crawl(url, depth, fetcher, msg, &wg, &s)
	go func() {
		wg.Wait()
		close(msg)
	}()
	for m := range msg {
		fmt.Println(m)
	}
}
func _crawl(url string, depth int, fetcher Fetcher, msg chan string, wg *sync.WaitGroup, s *SafeMap) {
	// TODO: catch url synchrously
	// TODO: no duplicated

	defer wg.Done()
	if depth <= 0 {
		return
	}
	if s.Contains(url) {
		return
	} else {
		s.Add(url)
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		msg <- err.Error()
		return
	}
	msg <- fmt.Sprintf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go _crawl(u, depth-1, fetcher, msg, wg, s)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFther
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is filled-up fakeFetcher
var fetcher = fakeFetcher{

	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
