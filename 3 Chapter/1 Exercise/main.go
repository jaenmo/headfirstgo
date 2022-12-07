package main

import ("fmt"
		"time")
// YOUR CODE HERE: Write a scoreSummary function that will
// work with the calls below.

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println("------------------------------------------------------")
	scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	scoreSummary("Jessie", 79.3, 99.1, 82.5)
	scoreSummary("Lamar", 82.2, 95.4, 77.6)

	var l float64 = 2.2
	var le int = 5
	fmt.Println(int(l) + le)

	ast := time.Now()
	year := ast.Year()
	fmt.Println(year)

	fmt.Printf("%.2f\n", 1.988)
	fmt.Printf("%.2f\n", 1.986)
	fmt.Printf("%.2f\n", 1.985)
	fmt.Printf("%.2f", 1.984)
}

func scoreSummary(a string, b float64, c float64, d float64) {
	e := (b + c + d) / 3
	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n",
	a, b, c, d, e)

121wwwavf
}
