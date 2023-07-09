package main

import "fmt"

// struct are like object in JS

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// 2nd way to declare
	// alex := person{"Alex", "Anderson"}
	var james person
	james.firstName = "James"
	james.lastName = "Bond"

	fmt.Printf("%+v", james)
	fmt.Println("")
	fmt.Println(james)
}
