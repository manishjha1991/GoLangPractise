package main

import "fmt"

func main() {
	x := 5
	y := 10
	if x < y {
		fmt.Printf("%d is less then %d\n", x, y)
	}

	// Switch
	color := "Red"
	switch color {
	case "Red":
		fmt.Println("color is not Red")
	case "Orange":
		fmt.Println("color is not Red")
	default:
		fmt.Println("Color is Green")
	}
}
