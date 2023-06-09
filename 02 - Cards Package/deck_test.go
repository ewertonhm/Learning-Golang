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
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.ReadDir(filename)

	d := newDeck()
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(filename)
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 5)
	if len(hand) != 5 {
		t.Errorf("Expected 5 cards in hand, got %v", len(hand))

	}
	if len(remainingDeck) != 11 {
		t.Errorf("Expected 11 cards in remainingDeck, got %v", len(remainingDeck))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()
	if d[0] == "Ace of Spades" && d[1] == "Two of Spades" && d[16] == "Four of Clubs" {
		t.Errorf("Shuffle not working, got d1 %v, d2 %v, d16 %v", d[0], d[1], d[16])
	}
}
