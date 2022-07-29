package main

import "fmt"

func main(){
	//if we don't use pointers, we don't modify the value
	amount := 6
	double(&amount)
	fmt.Println(amount)

	fmt.Println("###########")
	
	//We can have pointers of any type
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

	//Assign pointers to some values
	var myOnt int
	var myPointerInt *int
	myPointerInt = &myOnt
	fmt.Println(myPointerInt)

	var myPaint int
	myPainter := &myPaint
	fmt.Println(myPainter)

	//Short way to do it
	myTest := 6
	myTester := &myTest
	fmt.Println(myTester)
	fmt.Println(*myTester)

	//Exercise
	//var mayInteger *int
	mayInt := 5
	mayInteger := &mayInt
	fmt.Println(mayInteger)
	fmt.Println(*mayInteger)

	
fmt.Println("--------------------")
	var baba *float64 = createPointer()
	fmt.Println(baba) 
	fmt.Println(*baba) 
	*baba = 13.3
	fmt.Println(baba) 
	fmt.Println(*baba)



}
func double(number *int){
	*number *= 2
	fmt.Println(number)
}
func createPointer() *float64{
	myFloat := 18.5
	return &myFloat
}