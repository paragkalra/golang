package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	numbers := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	for _, suit := range suits {
		for _, number := range numbers {
			cards = append(cards, number+" of "+suit)
		}
	}
	return cards
}
