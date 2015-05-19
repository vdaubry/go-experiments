package main

import (
    "fmt"
)

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}


func CrawlRecursive(url string, depth int, fetcher Fetcher, quit chan int, done chan map[string] bool) {
    if depth<=0 {
        quit <- 0
        return
    }
    
    doneUrls := <- done
    hasIt, didIt := doneUrls[url]
    if didIt && hasIt {
        quit <- 0
        done <- doneUrls
        return
    } else {
        doneUrls[url] = true
    }
    done <- doneUrls
    
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        quit <- 0
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    childrenQuit := make(chan int)
    for _, u := range urls {
        go CrawlRecursive(u, depth-1, fetcher, childrenQuit, done)
    }
    for _ = range urls {
        <- childrenQuit
    }
    quit <- 0
}

func Crawl(url string, depth int, fetcher Fetcher) {
    done := make(chan map[string] bool, 1)
    done <- map[string]bool{url: false}
    quit := make(chan int)
    go CrawlRecursive(url, depth, fetcher, quit, done)
    <- quit //wait for quit
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
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