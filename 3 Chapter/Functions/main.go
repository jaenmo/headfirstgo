package main

import "fmt"

func main(){
	w1:= 4.2
	l1:= 3.0
	fmt.Println(area(w1, l1)/10, "liters needed")
	fmt.Printf("%.2f liters needed\n", area(w1, l1)/10)

	w2 := 5.2
	l2 := 3.5
	fmt.Println(area(w2, l2)/10, "liters needed")
	fmt.Printf("%.2f liters needed\n", area(w2, l2)/10)

}
func area(w, l float64) float64{
	return w * l
}