package main

func main() {

	//cards := newDeck()
	cards := newDeckFromFile("my_cards")
	cards.suffle()
	//cards.saveToFile("my_cards")
	cards.print()
}
