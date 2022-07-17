package main

import ("fmt"
	"time"
"strings")

func main (){

	var now time.Time = time.Now()
	var date int = now.Day()
	fmt.Println(date)
	
}