package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard() //calling the newCard function and assigning its return value to the card variable
	//card = "Five of Diamonds"

	fmt.Println(card)
}

//defining helper function that gives us the five of diamonds
func newCard() string {
	return "Five of Diamonds"
}
