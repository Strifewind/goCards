package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected a deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t * testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")


	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
