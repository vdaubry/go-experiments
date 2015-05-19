package main

import (
  "fmt"
  "time"
)

func say(s string, c chan string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
  c <- s
}

func main() {
  fmt.Println("begin")
  c := make(chan string)
  go say("hello", c)
  y := <-c
  fmt.Println("channel 1 : ", y)
  go say("world", c)
  y = <-c
  fmt.Println("channel 2 : ", y)
  fmt.Println("end")
}
