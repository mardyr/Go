package main

import (
	"testing"
	"os"
)

const DECK_SIZE = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()
	length := len(d)

	if length != DECK_SIZE {
		t.Errorf("Expected deck lenght of %d but got %d ", DECK_SIZE, length)
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	name := "_decktesting"
	os.Remove(name)

	deck := newDeck()
	deck.saveToFile(name)

	loadedDeck := newDeckFromFile(name)

	if len(loadedDeck) != DECK_SIZE {
		t.Errorf("Expected deck lenght of %d but got %d ", DECK_SIZE,len(loadedDeck))
	}
	os.Remove(name)
}
