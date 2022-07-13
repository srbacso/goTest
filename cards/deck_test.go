package main

import (
	"os"
	"testing"
)

const LENGTH = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != LENGTH {
		t.Errorf("Length should be %d and was found to be %d", LENGTH, len(d))
	}
	if d[0] != "Ace of Spades" || d[len(d)-1] != "King of Clubs" {
		t.Errorf("First and last cards not correct")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const FILENAME = "_deckTesting"
	if os.Remove(FILENAME) != nil {
	}
	d := newDeck()
	if d.saveToFile(FILENAME) != nil {
	}
	loadedDeck := newDeckFromFile(FILENAME)
	os.Remove(FILENAME)
	if len(loadedDeck) != LENGTH {
		t.Errorf("Unexpected length of deck: %v", len(loadedDeck))
	}

}
