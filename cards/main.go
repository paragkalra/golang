package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("#################Dealing cards##############################")
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("#################Remaining cards##############################")
	remainingCards.print()
	fmt.Println("#################Saving cards to file##############################")
	cards.saveToFile("mycards")
	fmt.Println("#################Reading cards from file##############################")
	newDeckFromFile("mycards").print()
}
