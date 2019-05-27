package main

import "fmt"

// Need To Read More About CLOSURES FUNCTION IN DETAILS

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()

	for index := 0; index < 10; index++ {
		fmt.Println(sum(index))
	}

}
