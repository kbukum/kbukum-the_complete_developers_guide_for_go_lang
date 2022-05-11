package main

import (
	"reflect"
	"testing"
)

/**
Test
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 16
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %d but got %d", expectedLength, len(d))
	}

	expectedFirstCard := "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card of deck is %s but got %s", expectedFirstCard, d[0])
	}

	expectedLastCard := "Four of Clubs"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card of deck is %s but got %s", expectedLastCard, d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := deck([]string{"Ace of Spades", "Two of Hearts", "Three of Diamonds", "Four of Clubs"})
	d1, d2 := deal(d, 2)
	d1.print()
	d2.print()
	expectedD1 := deck([]string{"Ace of Spades", "Two of Hearts"})
	expectedD2 := deck([]string{"Three of Diamonds", "Four of Clubs"})
	if !reflect.DeepEqual(d1, expectedD1) {
		t.Errorf("Expected d1 is %v but we got %v", expectedD1, d1)
	}
	if !reflect.DeepEqual(d2, expectedD2) {
		t.Errorf("Expected d2 is %v but we got %v", expectedD2, d2)
	}
}
