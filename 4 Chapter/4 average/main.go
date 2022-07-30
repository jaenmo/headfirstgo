package main

import "fmt"

func main(){
	meat := []float64{71.8, 56.2, 89.5}
	var sum float64

	for _, v := range meat{
		sum += v
		fmt.Println(sum/3)
	}
}