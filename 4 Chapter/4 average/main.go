package main

import "fmt"

func main(){
	meat := []float64{71.8, 56.2, 89.5}
	sum := 0.0

	for _, v := range meat{
		sum += v	
	}
	fmt.Println(sum)
	fmt.Println(len(meat))
	fmt.Println(sum/float64(len(meat)))


}