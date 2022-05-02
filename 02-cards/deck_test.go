package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("expected decked length of 16, got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("expected `Ace of Spades`, got: %v", deck[0])
	}
}
