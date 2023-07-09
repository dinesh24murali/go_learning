package main

import "fmt"

// struct are like object in JS

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // field name is optional
}

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "brown",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	// This a pointer shortcut in GO
	// Even though "updateLastName" is expecting "receiver" of pointer to person
	// Go will infer for us and convert person to pointer to person
	jim.updateLastName("white")

	// Update contact
	jim.contactInfo.email = "jimmy@mephone.io"

	jim.print()

	name := "Bill"

	fmt.Println(*&name)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}
