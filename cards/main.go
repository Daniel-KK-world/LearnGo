package main

func main() {

	cards := newDeck()

	//hand, remainingCards := deal(cards, 5)

	cards.print()

	//hand.print()
	//remainingCards.print()

	/*cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	now we call this from the deck.go file instead.
	for i, card := range cards {
		fmt.Println(i, card)
	}
	*/
}

/*
func newCard() string {
	return "Five of Diamonds"
}
*/
