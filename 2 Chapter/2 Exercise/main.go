package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter racer name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)

	fmt.Print("Enter racer rank: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	rank, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// YOUR CODE HERE: Write some code that looks at the
	// "rank" variable and sets the "medal" variable to an
	// appropriate string. If "rank" is 1, "medal" should
	// be set to "gold". If "rank" is 2, "medal" should be
	// "silver". A rank of 3 should get a "bronze" medal,
	// and any other rank should get a "participant" medal.
	var medal string
	
	if rank == 1{
		medal = "gold"
	} else if rank == 2{
		medal = "silver"
	}else if rank == 3{
		medal = "bronze"
	}else {
		medal = "participant"
	}

	fmt.Println(name, "gets a", medal, "medal!")
}
afes