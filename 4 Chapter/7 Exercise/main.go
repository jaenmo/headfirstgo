//check how many times a number is repeated in an aray

package main

import "fmt"

func main(){
	numbers := [100]int{4, 2, 6, 7, 8, 0, 1, 8, 7, 8,
		1, 5, 3, 2, 2, 1, 9, 6, 1, 0, 0, 0, 8, 4, 6,
		2, 2, 4, 7, 9, 6, 5, 9, 0, 5, 1, 1, 5, 4, 7,
		7, 9, 7, 8, 6, 3, 3, 3, 4, 8, 0, 4, 1, 1, 7,
		9, 8, 8, 1, 2, 3, 6, 4, 9, 2, 5, 8, 6, 7, 7,
		5, 4, 2, 9, 4, 4, 2, 2, 5, 5, 0, 0, 0, 9, 1,
		9, 5, 8, 0, 1, 1, 0, 5, 3, 8, 6, 3, 4, 4, 9}

	var ocurrences [10] int
	var sum int

	for _, number := range numbers{
		ocurrences[number]++
	}
	for number, count := range ocurrences{
		fmt.Println("The number", number, "is repeated", count, "times")
		sum += count
	}
	fmt.Println(sum)

}