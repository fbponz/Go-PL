package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var string_out string
	fmt.Println(" ")
	fmt.Println("Testing the joining strings with the library strings using join")

	startTime := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	stopTime := time.Now()

	fmt.Println("Time first duration:")
	elapsed := stopTime.Sub(startTime)
	fmt.Println(elapsed)

	var sep string
	fmt.Println(" ")
	fmt.Println("Testing the joining strings with loops and basic operations")
	startTimeTwo := time.Now()
	for index_number := 1; index_number < len(os.Args); index_number++ {
		string_out += sep + os.Args[index_number]
		sep = " "
	}
	fmt.Println(string_out)
	stopTimeTwo := time.Now()

	fmt.Println("Second first duration:")
	elapsed_t2 := stopTimeTwo.Sub(startTimeTwo)
	fmt.Println(elapsed_t2)
}
