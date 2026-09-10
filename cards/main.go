package main

func main() {

	//cards := newDeck()
	//cards.saveToFile("my_cards")

	//loading the deck from file storage
	cards := newDeckFromFile("my_cards")
	cards.print()

}
