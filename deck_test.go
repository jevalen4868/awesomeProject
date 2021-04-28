package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d:= newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20 but was %v", len(d))
	}

	expected := "Ace of Spades"
	if d[0] != expected {
		t.Errorf("Expected [%v] but was [%v]", expected, d[0])
	}

	expected = "Five of Clubs"
	if d[len(d) - 1] != expected {
		t.Errorf("Expected [%v] but was [%v]", expected, d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFile := "./_decktesting"
	os.Remove(testFile)

	d := newDeck()

	d.saveToFile(testFile)

	nd := newDeckFromFile(testFile)

	expected := 20
	if len(nd) != expected {
		t.Errorf("Expected deck length of %v but was %v", expected, len(nd))
	}
	os.Remove(testFile)
}