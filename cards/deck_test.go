package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	EXPECTED_SIZE := 20
	EXPECTED_FIRST_CARD := "Ace of Spades"
	EXPECTED_LAST_CARD := "King of Clubs"

	if deck.size() != EXPECTED_SIZE {
		t.Errorf("Expected deck length of %v, but got %v", EXPECTED_SIZE, len(deck))
	}

	if deck.nextCard() != EXPECTED_FIRST_CARD {
		t.Errorf("Expected first card %v, but got %v", EXPECTED_FIRST_CARD, deck.nextCard())
	}

	if deck.lastCard() != EXPECTED_LAST_CARD {
		t.Errorf("Expected last card %v, but got %v", EXPECTED_LAST_CARD, deck.lastCard())
	}
}

func TestBuildHand(t *testing.T) {
	deck := newDeck()
	hand := buildHand(deck, 3)

	EXPECTED_SIZE := 3
	EXPECTED_FIRST_CARD := "Ace of Spades"
	EXPECTED_LAST_CARD := "Jack of Spades"

	if hand.size() != EXPECTED_SIZE {
		t.Errorf("Expected hand length to be %v, but got %v", EXPECTED_SIZE, hand.size())
	}

	if hand.nextCard() != EXPECTED_FIRST_CARD {
		t.Errorf("Expected hand first card %v, but got %v", EXPECTED_FIRST_CARD, hand.nextCard())
	}

	if hand.lastCard() != EXPECTED_LAST_CARD {
		t.Errorf("Expected hand last card %v, but got %v", EXPECTED_LAST_CARD, hand.lastCard())
	}
}

func TestBuildRemainingDeck(t *testing.T) {
	deck := newDeck()
	remainingDeck := buildRemainingDeck(deck, 3)

	EXPECTED_SIZE := 17
	EXPECTED_FIRST_CARD := "Queen of Spades"
	EXPECTED_LAST_CARD := "King of Clubs"

	if remainingDeck.size() != EXPECTED_SIZE {
		t.Errorf("Expected remaining deck length to be %v, but got %v", EXPECTED_SIZE, remainingDeck.size())
	}

	if remainingDeck.nextCard() != EXPECTED_FIRST_CARD {
		t.Errorf("Expected remaining deck first card %v, but got %v", EXPECTED_FIRST_CARD, remainingDeck.nextCard())
	}

	if remainingDeck.lastCard() != EXPECTED_LAST_CARD {
		t.Errorf("Expected remaining deck last card %v, but got %v", EXPECTED_LAST_CARD, remainingDeck.lastCard())
	}
}

func TestDeal(t *testing.T) {
	deck := newDeck()
	handByDeal, reaminingDeckByDeal := deal(deck, 3)

	EXPECTED_HAND_SIZE := 3
	EXPECTED_REMAINING_DECK_SIZE := 17

	if handByDeal.size() != EXPECTED_HAND_SIZE {
		t.Errorf("Expected hand dy deal length to be %v, but got %v", EXPECTED_HAND_SIZE, handByDeal.size())
	}

	if reaminingDeckByDeal.size() != EXPECTED_REMAINING_DECK_SIZE {
		t.Errorf("Expected remaining deck by deal length to be %v, but got %v", EXPECTED_REMAINING_DECK_SIZE, reaminingDeckByDeal.size())
	}
}
