package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, But got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, But got %v", len(d))
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card of Five of Clubs, But got %v", len(d))
	}
}
