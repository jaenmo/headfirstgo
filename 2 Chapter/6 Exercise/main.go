
//6.If the player ra out of turns without guessing correctly:
//"Sorry. You didn't guess my number. It was: [target]"


package main

import (
    "fmt"
    "math/rand"
	"time"
)

func main() {
//1. Generate a random number from 1 to 100, and store it as a 
//target number for the player to guess.
seconds := time.Now().Unix()
rand.Seed(seconds)
target := rand.Intn(100) + 1
	fmt.Println(target)

//2. Prompt the player to guess what the target number is, 
//and store their response
	var guess int
	fmt.Println("I've chosen a number between 1 and 100")
	fmt.Println("can you guess it?")
	fmt.Scanln(&guess)

//3. If the player's guess is less than the target number, say 
//"Oops. Your guess was LOW". Else, "Oops. Your guess was HIGH"
	for try := 0; try < 10; try++{
		fmt.Printf("You have tried %v times", try)

	if guess < target {
		fmt.Println("Oops. Your guess was LOW")
	}else if guess > target{
		fmt.Println("Oops. Your guess was HIGH")	
	}else if guess == target{
		fmt.Println("Good job! You guessed it")
	}
}
//4. Allow the player to guess up to 10 times. Before each guess, 
//let them know how many guesses they have left

//5. If the player's guess is equal to the target: 
//"Good job! You guessed it". Then stop asking for new guesses.


}