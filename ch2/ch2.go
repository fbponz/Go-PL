package main

import "fmt"

func main() {
	var i int = 20
	var f float32 = float32(i)
	const ValueI = 20

	fmt.Println(i)
	fmt.Println(f)

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

	fmt.Println("These three variables has overflow")
}
