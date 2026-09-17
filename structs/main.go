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

	//fmt.Printf("%+v", jim) //print the entire struct with field names
	jimPointer := &jim             //create a pointer to the jim struct
	jimPointer.updateName("Jimmy") //now update the name using pointer receiver method
	jim.print()

}

//this method has a value receiver, so it will not modify the original struct
/*func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}*/

func (pointerToPerson *person) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
