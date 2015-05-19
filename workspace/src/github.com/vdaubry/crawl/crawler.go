package main

import (
  "fmt"
  "os"
  "strconv"
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

var maxOpenRequests, _ = strconv.Atoi(os.Getenv("MAXREQUESTS"))
var urlCount, _ = strconv.Atoi(os.Getenv("URLCOUNT"))
var requestSemaphore = make(chan int, maxOpenRequests)

type HttpResponse struct {
  url      string
  response *http.Response
  err      error
  totalTime time.Duration
}


func httpGet(url string, ch chan *HttpResponse) {
  requestSemaphore <- 1 // Block until put in the semaphore queue
  fmt.Printf("Fetching %s \n", url)
  start_time := time.Now().UTC()
  timeout := time.Duration(15 * time.Second)
  client := http.Client{
    Timeout: timeout,
  }
  resp, err := client.Get(url)
  ch <- &HttpResponse{url, resp, err, time.Since(start_time)}
  <- requestSemaphore // Dequeue from the semaphore
}

func asyncGetUrls(urls []string) (int, int) {
  ch := make(chan *HttpResponse)
  for _, url := range urls {
      go httpGet(url, ch)
  }
  
  successCount := 0
  failureCount := 0
  for response := range ch {
    if response.err == nil {
      fmt.Printf("%s status: %s in %s\n", response.url,
             response.response.Status, response.totalTime)
      successCount+=1
    } else {
      fmt.Printf("err: %s in %s\n", response.err, response.totalTime)
      failureCount+=1
    }
    
    if (successCount+failureCount) == len(urls) {
      close(ch)
    }
  }
  return successCount, failureCount
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
  
  return data["domains"][:urlCount]
}

func main() {
  var urls = readUrls("domains-fast.json")
  begin := time.Now().UTC()
  successCount, failureCount := asyncGetUrls(urls)
  fmt.Printf("Success : %d \n", successCount)
  fmt.Printf("Failure : %d \n", failureCount)
  fmt.Printf("Total : Get %d urls in %s \n", urlCount, time.Since(begin))
}