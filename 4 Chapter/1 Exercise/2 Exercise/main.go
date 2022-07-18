package main

import ("fmt"
		"strings")

func main(){
	broken := ("G# R#ck!")
	cambio := strings.NewReplacer("#", "o")
	arreglao := cambio.Replace(broken)
	fmt.Println(arreglao)
}