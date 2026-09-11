package main

import "testing"

// make a test for the newDeck function
func TestNewDeck(t *testing.T) {
	// so create a new deck, check the length and the first and last card

	d := newDeck()

	// chek that len is 16
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// check that deck[0] is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %v", d[0])
	}

	// check that deck[len(d) -1] is Four of clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the last position to be Four of Clubs, but got %v", d[len(d)-1])
	}
}
