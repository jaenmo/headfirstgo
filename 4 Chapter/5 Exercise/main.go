// create an array with the day of the week and print them with the index

package main

import "fmt"

func main(){
	week := [5]string{ "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	
	for i, v := range week{
		fmt.Println(i, v)
	}
	fmt.Println("-------------")
	for ind :=0 ; ind < len(week) ; ind++{
		fmt.Println(ind, week[ind])
	}

}