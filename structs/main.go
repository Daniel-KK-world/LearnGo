package main

import "fmt"

//struct to represent a person

type person struct {
	firstName string
	lastName  string
}

func main() {
	//first approach to creating a person
	barcaCoach := person{"Hansi", "Flick"}
	fmt.Println(barcaCoach)

	//second approach to creating a person
	jackie := person{
		firstName: "Jackie",
		lastName:  "Chan",
	}
	fmt.Println(jackie)

	//third approach to creating a person
	var cena person
	fmt.Printf("%+v", cena)

	//updating a struct
	cena.firstName = "John"
	cena.lastName = "Cena"
}

//next thing is embedding.
