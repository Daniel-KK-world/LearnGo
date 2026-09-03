package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	/* ...
	now we call this from the deck.go file instead.
	for i, card := range cards {
		fmt.Println(i, card)
	}
	*/

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
