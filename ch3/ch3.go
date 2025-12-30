package main

import "fmt"

func main() {
	fmt.Println("These three variables has overflow")
	var greetings []string = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	fmt.Println(greetings[:2])
	fmt.Println(greetings[1:4])
	fmt.Println(greetings[3:5])

	var message []rune = []rune("Hi and ")
	fmt.Println(message)

	type employee struct {
		firstName string
		lastName  string
		id        int
	}
	employee_list := employee{
		firstName: "Manolo",
		lastName:  "Revilla",
		id:        10,
	}

	fmt.Println(employee_list)
}
