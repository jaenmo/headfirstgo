package main

import "fmt"

func main(){
fmt.Println(average(2, 4, 3))
}
func average (numero ...float64)float64{
	
	var sum float64
	for _, v:= range numero{
		sum += v
	}
	size:= float64(len(numero))
	return sum/size
}