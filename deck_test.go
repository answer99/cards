package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first of Spades of Ace, but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected first of Clubs of Four, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testDeck")
	d := newDeck()
	d.saveToFile("_testDeck")

	deckFromFile := newDeckFromFile("_testDeck")

	if len(deckFromFile) != 16 {
		t.Errorf("Expected deck length of 16 read from file, but got %v", len(d))
	}

	os.Remove("_testDeck")
}
