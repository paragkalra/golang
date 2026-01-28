package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", "Two of Diamonds"}
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
