package main_test

import (
	. "github.com/vdaubry/square"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SquareRoot", func() {
  
  Describe ("new", func() {
    It("returns a square", func() {
      var square = NewSquareRoot(1.0)
      Expect(square.Number()).To(Equal(1.0))
    })
  })
  
  Describe ("SetNumber", func() {
    It("returns a square", func() {
      var square = NewSquareRoot(1.0)
      square.SetNumber(2.0)
      Expect(square.Number()).To(Equal(2.0))
    })
  })
  Describe ("Compute square root", func(){
    It("returns 2.0 for 4.0", func() {
      var square = NewSquareRoot(4.0)
      Expect(square.Value()).To(Equal(2.0))
    })
    
    It("returns 3.16 for 10", func() {
      var square = NewSquareRoot(10.0)
      Expect(square.Value()).To(Equal(3.162277660168379))
    })
    
    It("returns 31.63 for 1001.0", func() {
      var square = NewSquareRoot(1001.0)
      Expect(square.Value()).To(Equal(31.63858403911275))
    })
    
    It("returns 1004.36 for 1008751.0", func() {
      var square = NewSquareRoot(1008751.0)
      Expect(square.Value()).To(Equal(1004.3659691566615))
    })
  })
})
