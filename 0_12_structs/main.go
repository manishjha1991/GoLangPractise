package main

import (
	"fmt"
	"strconv"
)

// Creating Person Struct
type Person struct {
	// Long Hand
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// ShorHand

	// firstName, lastName, city, gender string
	// age                               int
}

// Greatting Method Value Receiver

func (p Person) greet() string {
	return "Hello My Name is " + p.firstName + " " + p.lastName + "and my age is " + "  " +
		strconv.Itoa(p.age)
}

// hasBirthDay Pointer Receiver

func (p *Person) hasBirthDay() {
	p.age++
}

// get married change last Name
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
	p.lastName = spouseLastName
}
func main() {
	// Initialisation

	Person1 := Person{firstName: "x", lastName: "Y", city: "Bangalore", gender: "f", age: 26}

	// AlterNative

	Person2 := Person{"Z", "P", "Delhi", "m", 30}

	// fmt.Println(Person1)
	// //Get Single Value
	// fmt.Println(Person1.firstName)
	Person1.hasBirthDay()
	Person1.getMarried("Kumar")
	Person2.getMarried("Kumar")
	fmt.Println(Person2.greet())
}
