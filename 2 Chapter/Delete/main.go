package main

import "fmt"

func main(){
	truth := true
	negate(&truth)
	fmt.Println(truth)

	lie := false
	negate(&lie)
	fmt.Println(lie)

}
func negate(myBoolean *bool) {
	*myBoolean =! *myBoolean
}

// func main(){
// amount := 6
// double(&amount)
// fmt.Println(amount)

// }
// func double(number *int){
// 	*number *= 2
// }