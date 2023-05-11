package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Wrong number of cards. Expected deck length of 52 cards, but got %v", len(d))
	}

	if d[0] != "A♠" { // A♠ is the 1st card of the deck and therefore the 1st element of the array.
		t.Errorf("Wrong first card. Expected A♠ but got %v", d[0])
	}

	if d[len(d)-1] != "K♥" { // K♣ is the last card of the deck and therefore the last element of the array.
		t.Errorf("Wrong last card. Expected K♣ but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
