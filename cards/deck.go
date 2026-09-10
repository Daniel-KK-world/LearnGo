package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type of 'deck' which is a slice of strings
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

// new function to loop through the deck and print out val of each card
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// new function to deal a hand of cards, we need a deck and hand size.
// we can return multiple values in Go.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//new function to convert deck to a string

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// savig to file storage
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

// reading from file storage
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) // returns a byte slice and an error
	if err != nil {
		//option #1 - log the error and return a call to newDeck()
		//option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}
