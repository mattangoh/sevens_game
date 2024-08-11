package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Rolls struct {
    roll1, roll2 int
    sum int
  }

type User struct {
  betAmount float64
  guess string
}

func main() {
  
  fmt.Println("**********")
  fmt.Println("Welcome to the Sevens Game!")
  fmt.Println("**********")
  
  // User inputs
  user := UserInput()
  fmt.Println("User Input Summary:")
  fmt.Printf("You bet $%v\n", user.betAmount)
  fmt.Printf("You guessed %v.\n", user.guess)
  

  // Dealer rolling interaction
  fmt.Println("Dealer is Rolling")

  for i := 0; i < 3; i++ {
   time.Sleep(1 * time.Second)
   fmt.Println(".")
  } 
  
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

func UserInput() User {
  var betAmount float64

  for {
    fmt.Println("Please enter your bet amount!")
    _, err := fmt.Scanln(&betAmount)
    
    if err != nil {
      fmt.Println("Invalid input! Please re-enter a valid number for your float")
      // In case of extra lines
      var discard string
      fmt.Scanln(&discard)
    } else if betAmount <= 0 {
        fmt.Println("Bet amount must be greater than $0")
    } else {
        break
    }
  }
   
  var guess string

  for {
    fmt.Println("What is your guess (over, under, sevens)")
    _, err := fmt.Scanln(&guess)

    if err != nil {
      fmt.Println("Invalid input for guess. Please input a valid string.")
      // In case of extra lines
      var discard string
      fmt.Scanln(&discard)
    } else if guess != "over" && guess != "under" && guess != "sevens" {
      fmt.Println("Invalid guess. Must be either 'over', 'under', or 'sevens'.")
    } else {
        break
    }
  }
  
  user := User {betAmount, guess}

  return user
}

