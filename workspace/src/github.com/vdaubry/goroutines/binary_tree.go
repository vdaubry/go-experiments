package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  walk_node(t, ch)
  close(ch)
}

func walk_node(t *tree.Tree, ch chan int) {
  if t != nil {
    ch <- t.Value
    walk_node(t.Left, ch)
    walk_node(t.Right, ch)
  }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  for i := range ch1 {
    for i != <-ch2 {
      return false
    }
  }
  return true 
}

func main() {
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for i := range ch {
      fmt.Println("i = ", i)
  }
  tree1 := tree.New(1)
  tree2 := tree.New(2)
  
  fmt.Println("Is identical = ", Same(tree1, tree1))
  fmt.Println("Is identical = ", Same(tree1, tree2))
}
