package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  var res = make(map[string]int)
  var words = strings.Split(s, " ")
  
  for i:=0; i<len(words); i++ {
    var word = words[i]
    res[word]=0
    for j:=0; j<len(words); j++ {
      if words[j] == word {
        res[word]+=1
      }
    }
    
  }
  return res
}

func main() {
  wc.Test(WordCount)
}
