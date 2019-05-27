package main

import "fmt"

func main() {
	//Define Map

	// emails := make(map[string]string)
	// //Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["mike"] = "mike@gmail.com"
	// emails["manish"] = "manish@gmail.com"

	//Short Hand

	emails := map[string]string{"Bob": "bob@gmail.com", "mike": "mike@gmail.com", "manish": "manish@gmail.com"}

	//Get KEY AND VALUE

	fmt.Println("Before Delete", emails)
	fmt.Println("Length", len(emails))
	fmt.Println(emails["BOB"])
	//DELETE
	delete(emails, "Bob")
	fmt.Println("After Delete", emails)

}
