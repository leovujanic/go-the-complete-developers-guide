package main

import (
	"os"
	"testing"
)

const expectedDeckLen int = 24

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != expectedDeckLen {
		t.Errorf("Expected deck length of %d, but got %v", expectedDeckLen, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Six of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != expectedDeckLen {
		t.Errorf("Expected deck length of %d, but got %v", expectedDeckLen, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
