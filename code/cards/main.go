package main

import (
	"fmt"
	"os"
)

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func saveHndToFile() {
	fileName := "my-cards.txt"
	handSize := 5
	cards := newDeck()
	cards.shuffle()

	hand, _ := deal(cards, handSize)

	handError := hand.saveToFile(fileName)

	if handError != nil {
		fmt.Println("Error: ", handError)
		os.Exit(1)
	}

	var deckFromFile = newDeckFromFile(fileName)
	deckFromFile.print()
}
