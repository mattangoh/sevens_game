# Sevens game (Go Implementation)

Here is my implementation of the Sevens game. It is a cute little gambling game whereby you place your bet amount beforehand, and guess whether the resulting value is "over", "under", or "sevens". The dealer then rolls two 6 sided dice. 

If the combined sum of the two dice is:
- Under 7, then this is considered "under"...
- Exactly 7, then this is considered "sevens"...
- Over 7, then this is considere "over"...

If you guess "under" and it is under 7, then you get 2x your bet value as your payout... Likewise for guessing "over" and the result being over 7. If you guess "sevens" and the result is 7, then you get 5x your bet value as your payout.

Needless to say, if you are not accurate in your guess, then you lose all of your initial bet!

That's it. All that's left is for you to clone this repo and run go build for your operating system.
