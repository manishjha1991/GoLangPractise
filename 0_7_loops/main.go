package main

import "fmt"

func main() {

	//Long Method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	//i = i + 1
	// 	i++

	// }
	//Short Method

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	//FizzBuzz Challange

	// Print Number Dividalbel by 3 and 5
	//Check If number is visiable by 15 if it is divisiable by 15 it will be dividalbe by 3 and 5 also

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("Number %d\n", i)
		}
	}
}
