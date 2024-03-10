package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// given
	d := newDeck()

	// when then
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestSaveToFile(t *testing.T) {
	// given
	const testFilename string = "_decktest"
	os.Remove(testFilename)
	deck := newDeck()

	// when
	deck.saveToFile(testFilename)

	loadedDeck := newDeckFromFile(testFilename)

	// then
	if loadedDeck == nil {
		t.Errorf("Expected deck with cards, but got an empty deck")
	} else if len(loadedDeck) != 16 {
		t.Errorf("Expected deck with 16 cards, but got %v", len(loadedDeck))
	}

	// clean up
	os.Remove(testFilename)
}
