package main
import "fmt"
func main(){
	ar := [3]int{1, 2, 3}

	fmt.Println(ar)
	ar[2] = 0
	fmt.Println(ar)

	notes := [5]int{11, 22, 33, 44, 55}
	fmt.Println(len(notes))

	for i:=0; i<(len(notes)); i++{
		fmt.Println(i, notes[i])
	}
	fmt.Println(len(ar))

	fmt.Println("-----------------")
	var sum int
	for _, value := range notes{
		
		sum += value
		fmt.Println(sum)
		fmt.Println(sum/5)

	}

}