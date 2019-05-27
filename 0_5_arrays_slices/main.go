package main

import "fmt"

func main() {

	// var frutArray [2]string
	// frutArray[0] = "Apple"
	// frutArray[1] = "Orange"

	//Declare and Assign Array

	// frutArray := [2]string{"Apple", "Orange"}

	// fmt.Println(frutArray)
	// fmt.Println(frutArray[0])

	//Slice Array

	frutSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(frutSlice))
	fmt.Println(frutSlice[1:3])
}
