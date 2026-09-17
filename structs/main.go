package main

import "fmt"

//struct to represent contact information
type contactInfo struct {
	email   string
	zipCode int
}

//struct to represent a person
type person struct {
	firstName   string
	lastName    string
	contactInfo //embedding contactInfo struct into person struct
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim.party@example.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim) //print the entire struct with field names
}
