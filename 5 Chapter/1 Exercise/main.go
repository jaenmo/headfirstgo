package main

import "fmt"

func main(){
	candidates := []string{"Grahan Bell","Grahan Bell", "Grahan Bell", "Grahan Bell",
	 "Brian Martin", "Brian Martin", "Brian Martin", "Brian Martin", "Brian Martin", "Brian Martin",
	  "Carlos Diaz", "Carlos Diaz", "Carlos Diaz", "Carlos Diaz"}

	//var ocurrences [] int
	//var sum int

	for i, name := range candidates{
		//ocurrences[i]++
		fmt.Println(i,name)

	}
	// for number, count := range ocurrences{
	// 	fmt.Println("The candidate", number, "is repeated", count, "times")
	// 	sum += count
	// }
	//fmt.Println(sum)
	rintln("niosss")
}