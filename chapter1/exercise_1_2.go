package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var string_out string
	for index_number := 1; index_number < len(os.Args); index_number++ {
		string_out = strconv.Itoa(index_number) + " - " + os.Args[index_number]
		fmt.Println(string_out)
	}
}
