package main

import "fmt"

func main(){

	/* local variable definition */
var a int = 100
var b int = 200

fmt.Printf("Before swap, value of a : %d\n", a )
fmt.Printf("Before swap, value of b : %d\n", b )

/* calling a function to swap the values */
fmt.Println(swapReference(&a, &b))
fmt.Println(swapValue(a, b))
fmt.Println(swapSimple(a, b))

fmt.Printf("After swap, value of a : %d\n", a )
fmt.Printf("After swap, value of b : %d\n", b )

}

func swapReference(x, y *int) int {  
	var tempR int
		tempR = *x 
		*x = *y
		*y = tempR
		return tempR
	}

func swapValue(x, y int) int{ 
	var tempV int
		tempV = x 
		x = y
		y = tempV
		return tempV
	}

func swapSimple(x, y int) (int, int){
	return y, x
}
