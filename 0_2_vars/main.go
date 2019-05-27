package main

import "fmt"

func main() {
	// Note

	/*
		int is an alias of Int32, so, int and Int32 are the same type. Int32 represents 32-bits (4-bytes) signed integer. Int32 occupies 32-bits (4-bytes) space in the memory. As per the 4-bytes data capacity, an Int32's value capacity is -2147483648 to +2147483647.

	*/
	var b int32 = 1
	var c = "Manish"
	//ShortHand
	// name := "Manish"
	// email := "Manish@gmail.com"
	//Shorthand
	name, email := "Manish", "Manish@gmail.com"
	fmt.Printf("B _______%T\n", b)
	fmt.Printf("C_______%T\n", c)
	fmt.Println(name, email, b, c)

}
