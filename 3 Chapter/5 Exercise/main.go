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

	fmt.Printf("myInt is %T\n", &myInt)
	fmt.Printf("mymyStr is %T\n", &myStr)
	fmt.Printf("myFloa is %T\n", &myFloa)
	fmt.Printf("myStru is %T\n", &myStru)
	fmt.Printf("myInter is %T\n", &myInter)

	var myOnt int
	var myPointerInt *int
	myPointerInt = &myOnt
	fmt.Println(myPointerInt)

	var myPaint int
	myPainter := &myPaint
	fmt.Println(myPainter)

}
func double(number int){
	number *= 2
	fmt.Println(number)
}