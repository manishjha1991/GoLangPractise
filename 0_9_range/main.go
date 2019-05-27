package main

import "fmt"

func main() {
	ids := []int{33, 34, 36, 37, 38}
	for i, id := range ids {
		fmt.Printf("%d ID %d\n", i, id)
	}
	//Not Using Index put _ on the place of index
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}
	//Add Together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("SUM %d\n", sum)

	// Map

	emails := map[string]string{"Bob": "bob@gmail.com", "mike": "mike@gmail.com", "manish": "manish@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s is key and %s is value\n", k, v)
	}
	//Only Key
	for k := range emails {
		fmt.Printf("%s is key \n", k)
	}

}
