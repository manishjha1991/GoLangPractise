package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a, b)
	// Check Type

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	// *Int  is memmory DataTypes

	// Use * to read val from address
	fmt.Printf("Read from memory address: %d\n", *b)

	// Change Value from Pointer

	*b = 20
	fmt.Println(a)
}
