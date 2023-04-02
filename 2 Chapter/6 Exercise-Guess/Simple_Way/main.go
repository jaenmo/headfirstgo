//Generate a random number from 1 to 100, and try to guess it.
//You have 10 attemps to guess it

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

seconds := time.Now().Unix()
rand.Seed(seconds)
target := rand.Intn(100) + 1
	fmt.Println(target)

	fmt.Println("I've chosen a number between 1 and 100")
	fmt.Println("can you guess it?")

	var guess int

	for try := 0; try < 10; try++{
		fmt.Scanln(&guess)

	if guess < target {
		fmt.Println("Oops. Your guess was LOW")
		
	}else if guess > target{
		fmt.Println("Oops. Your guess was HIGH")
			
	}else if guess == target{
		fmt.Println("Good job! You guessed it")
		break
	}
	fmt.Printf("You have %v more times\n", 9 - try)

}
fmt.Println("Sorry. You didn't guess my number. It was:", target)


}

