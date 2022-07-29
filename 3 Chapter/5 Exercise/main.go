package main

import "fmt"

func main(){
	amount := 6
	double(amount)
	fmt.Println(amount)

	var myInt *int
	var myStr *string
	var myFloa *float64
	var myStru *struct{}
	var myInter *interface{}



	fmt.Println(&myInt)
	fmt.Println(&myStr)
	fmt.Println(&myFloa)
	fmt.Println(&myStru)
	fmt.Println(&myInter)










	
}
func double(number int){
	number *= 2
	fmt.Println(number)
}