package main

import "fmt"

var originalCount, eatenCount int = 10, 4


func main (){
	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}