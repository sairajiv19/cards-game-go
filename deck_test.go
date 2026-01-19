package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length 16 got %v", len(d))
	} else {
		t.Logf("New Deck test executed sucessfully. Length -> %d", 16)
	}
}

func TestDeal(t *testing.T) {
	d1, d2 := newDeck().Deal(5)
	if len(d1)+len(d2) != 16 {
		t.Errorf("The resulting decks are missing elements, total size : %d", len(d1)+len(d2))
	} else {
		t.Log("Test Deal success")
	}
}

func TestSaveDeckandLoadDeck(t *testing.T) {
	os.Remove("check_decktesting.txt")

	d := newDeck()
	d.saveDeck("check_decktesting.txt")

	loadedDeck := loadDeck("check_decktesting.txt")
	if len(d) != len(loadedDeck) {
		t.Error("Missing Elements")
	}
	os.Remove("check_decktesting.txt")
}
