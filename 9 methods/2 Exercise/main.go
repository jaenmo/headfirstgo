package main

import "fmt"

type Number int

func (n *Number) Double (){
	*n *= 2
}

func main(){
	number:= Number (2)
	fmt.Println("value of Number before anything is", number)
	number.Double()
	fmt.Println("value of Number after is", number)
}