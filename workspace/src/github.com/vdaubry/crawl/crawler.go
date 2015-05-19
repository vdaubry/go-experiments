package main

import (
  "fmt"
  "net/http"
  "time"
  "io/ioutil"
  "encoding/json"
  "log"
)

var urls = []string{
  "http://www.rubyconf.com/",
  "http://golang.org/",
  "http://matt.aimonetti.net/",
}

type HttpResponse struct {
  url      string
  response *http.Response
  err      error
  totalTime time.Duration
}

func httpGet(url string, ch chan *HttpResponse) {
  fmt.Printf("Fetching %s \n", url)
  start_time := time.Now().UTC()
  timeout := time.Duration(15 * time.Second)
  client := http.Client{
    Timeout: timeout,
  }
  resp, err := client.Get(url)
  ch <- &HttpResponse{url, resp, err, time.Since(start_time)}
}

func asyncGetUrls(urls []string) {
  ch := make(chan *HttpResponse)
  for _, url := range urls {
      go httpGet(url, ch)
  }
  
  responseCount := 0
  for response := range ch {
    responseCount+=1
    if response.err == nil {
      fmt.Printf("%s status: %s in %s\n", response.url,
             response.response.Status, response.totalTime)
    } else {
      fmt.Printf("err: %s in %s\n", response.err, response.totalTime)
    }
    
    if responseCount == len(urls) {
      close(ch)
    }
    
  }
}

type domainUrls map[string][]string

func readUrls(filepath string) []string {
  var data domainUrls
  file, err := ioutil.ReadFile(filepath)
  if err != nil {
      log.Fatal(err)
  }
  err = json.Unmarshal(file, &data)
  if err != nil {
      log.Fatal(err)
  }
  
  return data["domains"][:70]
}

func main() {
  var urls = readUrls("domains-fast.json")
  asyncGetUrls(urls)
}