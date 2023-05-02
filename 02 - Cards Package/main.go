package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	readed_deck := newDeckFromFile("my_cards")
	readed_deck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
