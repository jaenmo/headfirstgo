package main

import "fmt"

type Liters float64
type Mililiters float64
type Gallon float64

func (l Liters) ToMililiters() Mililiters {
	return Mililiters(l * 1000)
}
func (m Mililiters) ToLiters() Liters {
	return Liters(m / 1000)
}
func main() {
	l := Liters(3)
	fmt.Printf("%.1f liters is %.1f mililiters\n", l, l.ToMililiters())
	ml := Mililiters(500)
	fmt.Printf("%.1f liters is %.1f mililiters", ml, ml.ToLiters())
}
