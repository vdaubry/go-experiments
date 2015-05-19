package main

import "fmt"
import "math"

type SquareRoot struct {
  number float64
}

func NewSquareRoot(number float64) *SquareRoot {
  return &SquareRoot{number}
}

func (s SquareRoot) Number() float64 {
  return s.number
}

func (s *SquareRoot) SetNumber(number float64) {
  s.number = number
}

func (s SquareRoot) Value() float64 {
  var z float64 = s.number
  var previous float64 = 0
  var iteration int = 0
  
  for (math.Abs(previous-z)) > 0.000000001 {
    previous = z
    z = z - (z*z-s.number)/(2.0*z)
    iteration += 1
  }
  fmt.Printf("Found result in %d iteration\n", iteration)
  return z
}