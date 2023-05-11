package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// New type -> Deck of cards
// Wich a slice of strings

type deck []string

func newDeck() deck {
	deckOfCards := deck{}

	cardSuits := []string{"♠", "♦", "♣", "♥"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			deckOfCards = append(deckOfCards, values+suits)
		}
	}
	return deckOfCards
}

func (d deck) print() {
	i := 0
	for _, card := range d {
		i++
		if i != 13 {
			fmt.Print(card, "\t")
		} else {
			fmt.Println(card)
			i = 0
		}
	}
	fmt.Println("")
	fmt.Println("")

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func (d deck) printFullDeck() {
// 	if len(d) == 52 {
// 		for _, card := range d {
// 			if string(card[0]) != "K" {
// 				fmt.Print(card, "\t")
// 			} else {
// 				fmt.Println(card)
// 			}
// 		}
// 	} else {
// 		fmt.Println("The deck of cards is not full!")
// 	}
// 	fmt.Println("")
// }

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 = Read/Write/Execute Permiss
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //?
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) shuffle() {
	souce := rand.NewSource(time.Now().UnixNano())
	r := rand.New(souce)
	for i := range d {
		newposition := r.Intn(len(d) - 1)
		// d[i] = d[newposition]
		// d[newposition] = d[i]
		d[i], d[newposition] = d[newposition], d[i] // swap values in two random positions in the deck.
	}
}
