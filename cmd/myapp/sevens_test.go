package main

import (
	"errors"
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

// Since I dont think I have the necessary skills yet to deal with user inputs and mocking that in the test code
// I opt to split my function rather than test it itself... Pros and Cons, but maybe something to revisit
// in the future.

func validateBet(amount float64) error { 
  if amount <=0 {
    return errors.New("Bet must be greater than 0.")
  }
  return nil
}

func TestUserInput_ValidateBet(t *testing.T) {
  if err := validateBet(50); err != nil {
    t.Errorf("Expected no error, but got %v", err)
  }

  if err := validateBet(-10); err == nil {
    t.Errorf("Expected error for negative bet amount, but got none")
  }

  if err := validateBet(0); err == nil {
    t.Errorf("Expected error for zero bet amount, but got none")
  }
}

func validateGuess(guess string) error {
  if guess != "over" && guess != "under" && guess != "sevens" {
    return errors.New("invalid guess, must be either 'over', 'under', or 'sevens'")
  }
  return nil
}

func TestUserInput_ValidateGuess(t *testing.T) {
  if err := validateGuess("over"); err != nil {
    t.Errorf("Expected no error, but got %v", err)
  }

  if err := validateGuess("invalid"); err == nil {
    t.Errorf("Expected an error for invalid guess, but got none")
  }
}
