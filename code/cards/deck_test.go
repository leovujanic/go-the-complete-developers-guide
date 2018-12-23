package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedDeckLen := 24

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
