package main

import "fmt"

// YOUR CODE HERE: Write a "perimeter" function.

// An 8.2 by 10 meter plot
// A 5 by 5.2 meter plot
// A 6.2 by 4.5 meter plot


func main() {
	// YOUR CODE HERE: Call "perimeter" three times.
	// Add the three return values together, and store the
	// total in the "total" variable.
	a, err := perimeter(8.2, 10)
	if err =!nil {
		log.Fatal(err)
	}
	b, err := perimeter(5, 5.2)
	c, err := perimeter(6.2, 4.5)
	total := a + b + c

	fmt.Printf("You'll need %.2f meters of fencing", total)
}
func perimeter (length, width float64) (float64, error){
	if length >0 && width >0{
	return (2*length) + (2*width), nil
	}else{
		return "can't print it"
	}
}