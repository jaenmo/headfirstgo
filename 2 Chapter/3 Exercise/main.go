package main

import ("fmt"
		"strings")

func main(){
	broken := ("G# R#ck!")
	cambio := strings.NewReplacer("#", "o")
	fixed := cambio.Replace(broken)
	fmt.Println(fixed)
}