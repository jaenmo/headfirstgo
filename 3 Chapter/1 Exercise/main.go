package main

import "fmt"

// YOUR CODE HERE: Write a scoreSummary function that will
// work with the calls below.


// func main() {
// 	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
// 		"Name", "Score 1", "Score 2", "Score 3", "Average")
// 	fmt.Println("------------------------------------------------------")
// 	scoreSummary("Jermaine", 95.4, 82.3, 74.6)
// 	scoreSummary("Jessie", 79.3, 99.1, 82.5)
// 	scoreSummary("Lamar", 82.2, 95.4, 77.6)
// }

	
// 	func scoreSummary(a string, b float64, c float64, d float64) {
// 		e := (b+c+d)/3
		
// 	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
// 		 a, b, c, d, e)

// 	}

func main(){
max(2, 5)
fmt.Println("oh")

}
func max(num1, num2 int) int {
 var result int

	if (num1 > num2) {
	   result = num1
	} else {
	   result = num2
	}
	return result 
 }
