package player

import (
	"testing"

	"blackjack/internal/deck"
)

// TestNewPlayer tests player creation
func TestNewPlayer(t *testing.T) {
	name := "Test Player"
	player := NewPlayer(name)

	if player.Name != name {
		t.Errorf("Expected player name %s, got %s", name, player.Name)
	}
	if player.State != Playing {
		t.Errorf("Expected initial state Playing, got %v", player.State)
	}
	if len(player.Hand) != 0 {
		t.Error("Expected empty hand")
	}
}

// TestAddCard tests adding cards to player's hand
func TestAddCard(t *testing.T) {
	player := NewPlayer("Test")

	// Create test cards
	card1, _ := deck.NewCard(deck.Hearts, deck.Ace)
	card2, _ := deck.NewCard(deck.Spades, deck.King)

	// Add first card
	player.AddCard(card1)
	if len(player.Hand) != 1 {
		t.Errorf("Expected hand size 1, got %d", len(player.Hand))
	}

	// Add second card (should be blackjack)
	player.AddCard(card2)
	if len(player.Hand) != 2 {
		t.Errorf("Expected hand size 2, got %d", len(player.Hand))
	}
	if player.State != BlackJack {
		t.Errorf("Expected state BlackJack, got %v", player.State)
	}

	// Create a new player for bust test
	bustPlayer := NewPlayer("Bust Test")

	// Add cards that will cause a bust
	card3, _ := deck.NewCard(deck.Hearts, deck.King)
	card4, _ := deck.NewCard(deck.Diamonds, deck.Queen)
	card5, _ := deck.NewCard(deck.Clubs, deck.Jack)

	bustPlayer.AddCard(card3) // 10
	bustPlayer.AddCard(card4) // 20
	bustPlayer.AddCard(card5) // 30 (bust)

	if bustPlayer.State != Busted {
		t.Errorf("Expected state Busted, got %v", bustPlayer.State)
	}
}

// TestGetHandValue tests hand value calculation
func TestGetHandValue(t *testing.T) {
	tests := []struct {
		name     string
		cards    []deck.Card
		expected int
	}{
		{
			name: "Blackjack",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.King),
			},
			expected: 21,
		},
		{
			name: "Multiple Aces",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.Ace),
				mustCreateCard(t, deck.Diamonds, deck.Nine),
			},
			expected: 21,
		},
		{
			name: "Simple Hand",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Five),
				mustCreateCard(t, deck.Spades, deck.Ten),
			},
			expected: 15,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			player := NewPlayer("Test")
			for _, card := range test.cards {
				player.AddCard(card)
			}

			got := player.GetHandValue()
			if got != test.expected {
				t.Errorf("Expected hand value %d, got %d", test.expected, got)
			}
		})
	}
}

// TestStand tests the stand functionality
func TestStand(t *testing.T) {
	player := NewPlayer("Test")
	player.Stand()

	if player.State != Standing {
		t.Errorf("Expected state Standing, got %v", player.State)
	}
}

// TestClearHand tests clearing the player's hand
func TestClearHand(t *testing.T) {
	player := NewPlayer("Test")

	// Add some cards
	card, _ := deck.NewCard(deck.Hearts, deck.Ace)
	player.AddCard(card)

	// Clear hand
	player.ClearHand()

	if len(player.Hand) != 0 {
		t.Error("Expected empty hand after clear")
	}
	if player.State != Playing {
		t.Error("Expected state to reset to Playing")
	}
}

// TestHasBlackjack tests blackjack detection
func TestHasBlackjack(t *testing.T) {
	tests := []struct {
		name     string
		cards    []deck.Card
		expected bool
	}{
		{
			name: "Has Blackjack",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.King),
			},
			expected: true,
		},
		{
			name: "21 with three cards",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Seven),
				mustCreateCard(t, deck.Spades, deck.Seven),
				mustCreateCard(t, deck.Diamonds, deck.Seven),
			},
			expected: false,
		},
		{
			name: "Not Blackjack",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ten),
				mustCreateCard(t, deck.Spades, deck.Nine),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			player := NewPlayer("Test")
			for _, card := range test.cards {
				player.AddCard(card)
			}

			got := player.HasBlackjack()
			if got != test.expected {
				t.Errorf("HasBlackjack() = %v, want %v", got, test.expected)
			}
		})
	}
}

// Helper function to create cards for testing
func mustCreateCard(t *testing.T, suit deck.Suit, rank deck.Rank) deck.Card {
	card, err := deck.NewCard(suit, rank)
	if err != nil {
		t.Fatalf("Failed to create card: %v", err)
	}
	return card
}
