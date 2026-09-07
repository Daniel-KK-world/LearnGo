package main

import "fmt"

//create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//new function to loop through the deck and print out val of each card
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//new function to deal a hand of cards, we need a deck and hand size.
//we can return multiple values in Go.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
