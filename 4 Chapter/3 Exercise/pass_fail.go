// we need to write a program that allows a student to type in their percentage grade and tells
//  them whether they passed or not. Passing or failing follows a simple formula: a grade of 60%
//   or more is passing, and less than 60% is failing. So our program will need to give one response 
// if the percentage users enter is 60 or greater, and a different response otherwise

package main

import "fmt"

func main(){
	fmt.Println("Type your percentage grade")
	var grade int
	fmt.Scanln(&grade)

	if grade < 60 {
		fmt.Println("You have failed :(")
	}else {
		fmt.Println("You have passed! :)")
	}

}