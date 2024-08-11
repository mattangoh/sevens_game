package main

import (
	"fmt"
	"math/rand"
)

type Rolls struct {
    roll1, roll2 int
    sum int
  }

func main() {
  
  fmt.Println("**********")
  fmt.Println("Welcome to the Sevens Game!")
  fmt.Println("**********")

  dealer := DealerRoll()

  fmt.Printf("The first value is %v. The second value is %v, the sum is %v",
  dealer.roll1, dealer.roll2, dealer.sum)

}

func DealerRoll() Rolls {
  roll1 := rand.Intn(6) + 1 
  roll2 := rand.Intn(6) + 1
  sum := roll1 + roll2

  dealer := Rolls{roll1, roll2, sum}
  
  return dealer
}
