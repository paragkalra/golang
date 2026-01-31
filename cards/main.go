package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println("###############################################")
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("###############################################")
	remainingCards.print()
}
