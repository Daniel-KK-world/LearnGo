package main

import "fmt"

//create a new type of 'deck' which is a slice of strings
type deck []string

//new function to loop through the deck and print out val of each card
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
