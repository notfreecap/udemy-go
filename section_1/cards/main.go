package main

func main() {
	const filename string = "static_deck.text"
	deck := newDeck()
	deck.shuffle()

	hand, deck := deck.deal(4)

	hand.print()
	deck.print()
	hand.saveToFile(filename)

	newDeckFromFile(filename).print()
}
