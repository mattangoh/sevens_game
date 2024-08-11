package main

import (
  "testing"
)

func TestDealerRoll(t *testing.T) {
  
  //total simulations
  numRolls := 100

  for i := 0; i < numRolls; i++ {
    dealer := DealerRoll()

    // Check sum of x and y is accurate
    if dealer.sum != dealer.roll1 + dealer.roll2 {
      t.Errorf("Expected sum %v, but got %v", dealer.roll1+dealer.roll2, dealer.sum)
    }
    
    // Check roll1 falls between 1 and 6
    if dealer.roll1 < 1 || dealer.roll1 > 6 {
      t.Errorf("Expected roll1 to be between 1 and 6, but got %v", dealer.roll1)
    }

    // Check roll2 falls between 1 and 6
    if dealer.roll2 < 1 || dealer.roll2 > 6 {
      t.Errorf("Expected roll2 to be between 1 and 6, but got %v", dealer.roll2)
    }
  }
}
