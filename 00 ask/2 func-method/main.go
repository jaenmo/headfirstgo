package main

import ("fmt"
		"math")

type circle struct{
	radius float64
}
func(c circle) area() float64{
	return math.Pi * c.radius * c.radius
}
func main(){
	x := circle{2.5}
	fmt.Println(x.area())

	circle := circle{radius: 3}
	fmt.Println(circle.area())


}