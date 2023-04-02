package main

import "fmt"

func main() {
	fmt.Println(maximum(1, 3.2, 6.4, 2.2))
}
func maximum(numbers ...float64) float64 {
	var max float64
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}
