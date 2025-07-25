package deck

import (
	"strings"
	"testing"
)

// TestNewDeck tests deck creation and initial state
func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	// Test deck size
	if deck.RemainingCards() != 52 {
		t.Errorf("Expected deck size of 52, got %d", deck.RemainingCards())
	}

	// Count occurrences of each suit and rank
	suitCount := make(map[Suit]int)
	rankCount := make(map[Rank]int)
	seenCards := make(map[string]bool) // To check for duplicates

	for _, card := range deck.cards {
		suitCount[card.Suit]++
		rankCount[card.Rank]++

		// Check for duplicate cards
		cardKey := card.String()
		if seenCards[cardKey] {
			t.Errorf("Found duplicate card: %s", cardKey)
		}
		seenCards[cardKey] = true
	}

	// Check suit distribution
	for suit, count := range suitCount {
		if count != 13 {
			t.Errorf("Expected 13 cards of suit %s, got %d", suit, count)
		}
	}

	// Check rank distribution
	for rank, count := range rankCount {
		if count != 4 {
			t.Errorf("Expected 4 cards of rank %s, got %d", rank, count)
		}
	}
}

// TestShuffle tests deck shuffling
func TestShuffle(t *testing.T) {
	// Test with full deck
	t.Run("Full deck shuffle", func(t *testing.T) {
		deck1 := NewDeck()
		deck2 := NewDeck()

		// Store initial order
		initialOrder := make([]Card, len(deck1.cards))
		copy(initialOrder, deck1.cards)

		// Shuffle first deck
		deck1.Shuffle()

		// Count differences
		differences := 0
		for i := 0; i < 52; i++ {
			if deck1.cards[i] != deck2.cards[i] {
				differences++
			}
		}

		// It's extremely unlikely that less than 20 cards would be in different positions
		if differences < 20 {
			t.Errorf("Shuffle didn't appear to randomize the deck well. Only %d cards in different positions", differences)
		}

		// Check that no cards were lost or added
		if len(deck1.cards) != 52 {
			t.Errorf("Expected 52 cards after shuffle, got %d", len(deck1.cards))
		}

		// Check that all original cards are still present
		presentCards := make(map[string]bool)
		for _, card := range deck1.cards {
			presentCards[card.String()] = true
		}
		for _, card := range initialOrder {
			if !presentCards[card.String()] {
				t.Errorf("Card %s missing after shuffle", card.String())
			}
		}
	})

	// Test with empty deck
	t.Run("Empty deck shuffle", func(t *testing.T) {
		deck := NewDeck()
		// Draw all cards
		for i := 0; i < 52; i++ {
			deck.DrawCard()
		}
		// Shuffle empty deck (shouldn't panic)
		deck.Shuffle()
		if len(deck.cards) != 0 {
			t.Error("Empty deck should remain empty after shuffle")
		}
	})

	// Test with single card
	t.Run("Single card shuffle", func(t *testing.T) {
		deck := NewDeck()
		// Draw all but one card
		for i := 0; i < 51; i++ {
			deck.DrawCard()
		}
		// Shuffle with single card (shouldn't panic)
		deck.Shuffle()
		if len(deck.cards) != 1 {
			t.Error("Expected 1 card after shuffling single card deck")
		}
	})
}

// TestDrawCard tests drawing cards from the deck
func TestDrawCard(t *testing.T) {
	t.Run("Draw all cards", func(t *testing.T) {
		deck := NewDeck()
		drawnCards := make(map[string]bool)

		// Draw all 52 cards
		for i := 0; i < 52; i++ {
			card, err := deck.DrawCard()
			if err != nil {
				t.Errorf("Unexpected error drawing card %d: %v", i+1, err)
			}

			// Check for duplicate draws
			cardKey := card.String()
			if drawnCards[cardKey] {
				t.Errorf("Drew duplicate card: %s", cardKey)
			}
			drawnCards[cardKey] = true
		}

		// Verify deck is empty
		if deck.RemainingCards() != 0 {
			t.Errorf("Expected 0 cards remaining, got %d", deck.RemainingCards())
		}

		// Try to draw from empty deck
		_, err := deck.DrawCard()
		if err == nil {
			t.Error("Expected error drawing from empty deck")
		}
	})

	t.Run("Draw and verify order", func(t *testing.T) {
		deck := NewDeck()
		firstCard := deck.cards[0]
		drawnCard, err := deck.DrawCard()

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if drawnCard != firstCard {
			t.Error("Drew different card than expected")
		}
		if deck.RemainingCards() != 51 {
			t.Errorf("Expected 51 cards remaining, got %d", deck.RemainingCards())
		}
	})
}

// TestRemainingCards tests the RemainingCards method
func TestRemainingCards(t *testing.T) {
	tests := []struct {
		name      string
		drawCount int
		expected  int
	}{
		{"New deck", 0, 52},
		{"After one draw", 1, 51},
		{"After ten draws", 10, 42},
		{"Empty deck", 52, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			deck := NewDeck()
			for i := 0; i < test.drawCount; i++ {
				deck.DrawCard()
			}
			remaining := deck.RemainingCards()
			if remaining != test.expected {
				t.Errorf("Expected %d cards remaining, got %d", test.expected, remaining)
			}
		})
	}
}

// TestString tests the String method
func TestString(t *testing.T) {
	t.Run("Full deck", func(t *testing.T) {
		deck := NewDeck()
		str := deck.String()

		// Check basic format
		if !strings.Contains(str, "Deck with 52 cards") {
			t.Error("String should contain deck size")
		}

		// Check that all cards are listed
		for _, card := range deck.cards {
			if !strings.Contains(str, card.String()) {
				t.Errorf("String missing card: %s", card.String())
			}
		}
	})

	t.Run("Empty deck", func(t *testing.T) {
		deck := NewDeck()
		// Draw all cards
		for i := 0; i < 52; i++ {
			deck.DrawCard()
		}

		str := deck.String()
		if !strings.Contains(str, "Deck with 0 cards") {
			t.Error("String should show empty deck")
		}
	})

	t.Run("Partial deck", func(t *testing.T) {
		deck := NewDeck()
		// Draw some cards
		for i := 0; i < 10; i++ {
			deck.DrawCard()
		}

		str := deck.String()
		if !strings.Contains(str, "Deck with 42 cards") {
			t.Error("String should show correct number of remaining cards")
		}
	})
}
