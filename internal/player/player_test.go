package player

import (
	"strings"
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

// TestString tests the String method
func TestString(t *testing.T) {
	player := NewPlayer("Test Player")

	// Test empty hand
	str := player.String()
	if !strings.Contains(str, "Test Player") || !strings.Contains(str, "Value: 0") {
		t.Error("String representation should contain player name and initial value")
	}

	// Test with cards
	card, _ := deck.NewCard(deck.Hearts, deck.Ace)
	player.AddCard(card)
	str = player.String()
	if !strings.Contains(str, "Ace of Hearts") {
		t.Error("String representation should contain card description")
	}
}

// TestStateTransitions tests all possible state transitions
func TestStateTransitions(t *testing.T) {
	tests := []struct {
		name          string
		actions       func(*Player)
		expectedState PlayerState
		expectedValue int
	}{
		{
			name:          "Initial state",
			actions:       func(p *Player) {},
			expectedState: Playing,
			expectedValue: 0,
		},
		{
			name: "Stand from Playing",
			actions: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ten))
				p.Stand()
			},
			expectedState: Standing,
			expectedValue: 10,
		},
		{
			name: "Bust from Playing",
			actions: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ten))
				p.AddCard(mustCreateCard(t, deck.Spades, deck.Ten))
				p.AddCard(mustCreateCard(t, deck.Diamonds, deck.Two)) // 22 total
			},
			expectedState: Busted,
			expectedValue: 22,
		},
		{
			name: "BlackJack from Playing",
			actions: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ace))
				p.AddCard(mustCreateCard(t, deck.Spades, deck.King))
			},
			expectedState: BlackJack,
			expectedValue: 21,
		},
		{
			name: "Stay Playing under 21",
			actions: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ten))
				p.AddCard(mustCreateCard(t, deck.Spades, deck.Five))
			},
			expectedState: Playing,
			expectedValue: 15,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			player := NewPlayer("Test")
			test.actions(player)

			if player.State != test.expectedState {
				t.Errorf("Expected state %v, got %v", test.expectedState, player.State)
			}
			if player.GetHandValue() != test.expectedValue {
				t.Errorf("Expected hand value %d, got %d", test.expectedValue, player.GetHandValue())
			}
		})
	}
}

// TestGetHandValueWithAces tests complex Ace value calculations
func TestGetHandValueWithAces(t *testing.T) {
	tests := []struct {
		name     string
		cards    []deck.Card
		expected int
	}{
		{
			name: "Single Ace",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
			},
			expected: 11,
		},
		{
			name: "Two Aces",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.Ace),
			},
			expected: 12, // One Ace becomes 1
		},
		{
			name: "Three Aces",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.Ace),
				mustCreateCard(t, deck.Diamonds, deck.Ace),
			},
			expected: 13, // Two Aces become 1
		},
		{
			name: "Ace with Face Card",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.King),
			},
			expected: 21,
		},
		{
			name: "Ace becomes 1 to prevent bust",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.Ten),
				mustCreateCard(t, deck.Diamonds, deck.Five),
			},
			expected: 16, // Ace must be 1 to prevent bust
		},
		{
			name: "Multiple Aces with high cards",
			cards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.Ace),
				mustCreateCard(t, deck.Diamonds, deck.King),
			},
			expected: 12, // Both Aces must be 1
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

// TestPlayerString tests all string representation scenarios
func TestPlayerString(t *testing.T) {
	tests := []struct {
		name          string
		setupPlayer   func(*Player)
		expectedParts []string // Strings that should be present
	}{
		{
			name:        "New player",
			setupPlayer: func(p *Player) {},
			expectedParts: []string{
				"Player: Test",
				"Hand:",
				"Value: 0",
				"State: Playing",
			},
		},
		{
			name: "Player with cards",
			setupPlayer: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ace))
				p.AddCard(mustCreateCard(t, deck.Spades, deck.King))
			},
			expectedParts: []string{
				"Player: Test",
				"Ace of Hearts",
				"King of Spades",
				"Value: 21",
				"State: BlackJack",
			},
		},
		{
			name: "Busted player",
			setupPlayer: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.King))
				p.AddCard(mustCreateCard(t, deck.Spades, deck.Queen))
				p.AddCard(mustCreateCard(t, deck.Diamonds, deck.Jack))
			},
			expectedParts: []string{
				"Value: 30",
				"State: Busted",
			},
		},
		{
			name: "Standing player",
			setupPlayer: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ten))
				p.Stand()
			},
			expectedParts: []string{
				"Value: 10",
				"State: Standing",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			player := NewPlayer("Test")
			test.setupPlayer(player)

			str := player.String()
			for _, expectedPart := range test.expectedParts {
				if !strings.Contains(str, expectedPart) {
					t.Errorf("Expected string to contain %q, got %q", expectedPart, str)
				}
			}
		})
	}
}

// TestClearHandAfterStates tests clearing hand from different states
func TestClearHandAfterStates(t *testing.T) {
	tests := []struct {
		name       string
		setupState func(*Player)
	}{
		{
			name: "Clear after BlackJack",
			setupState: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ace))
				p.AddCard(mustCreateCard(t, deck.Spades, deck.King))
			},
		},
		{
			name: "Clear after Bust",
			setupState: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.King))
				p.AddCard(mustCreateCard(t, deck.Spades, deck.Queen))
				p.AddCard(mustCreateCard(t, deck.Diamonds, deck.Jack))
			},
		},
		{
			name: "Clear after Standing",
			setupState: func(p *Player) {
				p.AddCard(mustCreateCard(t, deck.Hearts, deck.Ten))
				p.Stand()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			player := NewPlayer("Test")
			test.setupState(player)
			player.ClearHand()

			if player.State != Playing {
				t.Errorf("Expected state Playing after clear, got %v", player.State)
			}
			if len(player.Hand) != 0 {
				t.Error("Expected empty hand after clear")
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
