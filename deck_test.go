package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length to be 16, but got %d", len(d))
	}

	if d[0] != "One of Spades" {
		t.Errorf("Expected first element to be 'One Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "Ace of Diamonds" {
		t.Errorf("Expected first element to be 'Ace of Diamonds', but got '%v'", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	clean()

	d := newDeck()
	d.saveToFile("_decktesting")

	deskFromFile := newDeckFromFile("_decktesting")
	if len(deskFromFile) != 16 {
		t.Errorf("Expected deck length to be 16, but got %d", len(deskFromFile))
	}

	clean()
}

func clean() {
	os.Remove("_decktesting")
}
