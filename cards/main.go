package main

import "fmt"

func main() {
	// card := newCard()
	// fmt.Println(card)

	// cards := deck{"A♥", "2♥", "3♥", "4♥", "5♥", "6♥", "7♥", "8♥"}
	// fmt.Println(cards)
	// cards = append(cards, newCard())
	// fmt.Println(cards)
	// cards = append(cards, "2♦", "3♦")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()
	// fmt.Println(" ------------ ")

	// cards := newDeck()
	// // cards.print()
	// // fmt.Println("")
	// hand, remainigCards := deal(cards, 4)
	// hand.print()
	// remainigCards.print()
	// fmt.Println(" ------------ ")
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	//fmt.Println(cards)
	fmt.Println("Deck of cards from file: ")
	cards.print()
	fmt.Println(" ------------ ")

	cards.shuffle()
	fmt.Println("Deck of cards shuffled: ")
	cards.print()
	fmt.Println(" ------------ ")

}
