// Check how many times is repeated one number in an array

package main

import "fmt"

func main() {
	// We'll count the number of times each number occurs
	// within this array.
	numbers := [9]int{1, 0, 2, 0, 1, 0, 0, 1, 2}
	// YOUR CODE HERE: Process each element of "numbers",
	// keeping track of how many times you've seen 0, 1,
	// or 2.
	// Finally, print out how many times each number
	// occurred.

	ocurrences := [3]int{}

	for _, number := range numbers{
		ocurrences [number] ++
	}
	for number, count := range ocurrences{
		fmt.Println(number, count)
	}
}




