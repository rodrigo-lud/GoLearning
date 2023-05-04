package main

import "fmt"

func main() {
	// card := newCard()
	// fmt.Println(card)

	cards := []string{"A♥", "2♥", "3♥", "4♥", "5♥", "6♥", "7♥", "8♥"}
	fmt.Println(cards)
	cards = append(cards, newCard())
	fmt.Println(cards)
	cards = append(cards, "2♦", "3♦")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "A♦"
}
