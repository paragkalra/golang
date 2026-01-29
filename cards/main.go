package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Five of Diamonds"
	card := newCard()
	fmt.Println(card)
	cards := deck{"Ace of Diamonds", "Two of Diamonds", newCard()}
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
