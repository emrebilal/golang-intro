package main

func main() {
	cards := newDeck()

	// Print of all
	//cards.print()

	// Assigned to two variable because deal() returns two values
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()

	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")

	//c := newDeckFromFile("my_cards")
	//c.print()

	cards.shuffle()
	cards.print()
}
