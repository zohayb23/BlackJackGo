package deck

import (
	"testing"
)

// TestNewCard tests the creation of cards
func TestNewCard(t *testing.T) {
	tests := []struct {
		name        string
		suit        Suit
		rank        Rank
		expectError bool
	}{
		{
			name:        "Valid card - Ace of Hearts",
			suit:        Hearts,
			rank:        Ace,
			expectError: false,
		},
		{
			name:        "Valid card - Ten of Diamonds",
			suit:        Diamonds,
			rank:        Ten,
			expectError: false,
		},
		{
			name:        "Invalid suit",
			suit:        "InvalidSuit",
			rank:        Ace,
			expectError: true,
		},
		{
			name:        "Invalid rank",
			suit:        Hearts,
			rank:        "InvalidRank",
			expectError: true,
		},
		{
			name:        "Empty suit",
			suit:        "",
			rank:        Ace,
			expectError: true,
		},
		{
			name:        "Empty rank",
			suit:        Hearts,
			rank:        "",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			card, err := NewCard(test.suit, test.rank)

			if test.expectError {
				if err == nil {
					t.Error("Expected an error but didn't get one")
				}
			} else {
				if err != nil {
					t.Errorf("Didn't expect an error but got: %v", err)
				}
				if card.Suit != test.suit {
					t.Errorf("Expected suit %v, got %v", test.suit, card.Suit)
				}
				if card.Rank != test.rank {
					t.Errorf("Expected rank %v, got %v", test.rank, card.Rank)
				}
			}
		})
	}
}

// TestCardValue tests all possible card values
func TestCardValue(t *testing.T) {
	tests := []struct {
		name     string
		card     Card
		expected int
	}{
		{
			name:     "Ace value",
			card:     Card{Suit: Hearts, Rank: Ace},
			expected: 11,
		},
		{
			name:     "King value",
			card:     Card{Suit: Hearts, Rank: King},
			expected: 10,
		},
		{
			name:     "Queen value",
			card:     Card{Suit: Hearts, Rank: Queen},
			expected: 10,
		},
		{
			name:     "Jack value",
			card:     Card{Suit: Hearts, Rank: Jack},
			expected: 10,
		},
		{
			name:     "Ten value",
			card:     Card{Suit: Hearts, Rank: Ten},
			expected: 10,
		},
		{
			name:     "Nine value",
			card:     Card{Suit: Hearts, Rank: Nine},
			expected: 9,
		},
		{
			name:     "Eight value",
			card:     Card{Suit: Hearts, Rank: Eight},
			expected: 8,
		},
		{
			name:     "Seven value",
			card:     Card{Suit: Hearts, Rank: Seven},
			expected: 7,
		},
		{
			name:     "Six value",
			card:     Card{Suit: Hearts, Rank: Six},
			expected: 6,
		},
		{
			name:     "Five value",
			card:     Card{Suit: Hearts, Rank: Five},
			expected: 5,
		},
		{
			name:     "Four value",
			card:     Card{Suit: Hearts, Rank: Four},
			expected: 4,
		},
		{
			name:     "Three value",
			card:     Card{Suit: Hearts, Rank: Three},
			expected: 3,
		},
		{
			name:     "Two value",
			card:     Card{Suit: Hearts, Rank: Two},
			expected: 2,
		},
		{
			name:     "Invalid rank value",
			card:     Card{Suit: Hearts, Rank: "InvalidRank"},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.card.Value()
			if got != test.expected {
				t.Errorf("Card.Value() = %v, want %v", got, test.expected)
			}
		})
	}
}

// TestIsFaceCard tests face card detection
func TestIsFaceCard(t *testing.T) {
	tests := []struct {
		name     string
		card     Card
		expected bool
	}{
		{
			name:     "King is face card",
			card:     Card{Suit: Hearts, Rank: King},
			expected: true,
		},
		{
			name:     "Queen is face card",
			card:     Card{Suit: Hearts, Rank: Queen},
			expected: true,
		},
		{
			name:     "Jack is face card",
			card:     Card{Suit: Hearts, Rank: Jack},
			expected: true,
		},
		{
			name:     "Ace is not face card",
			card:     Card{Suit: Hearts, Rank: Ace},
			expected: false,
		},
		{
			name:     "Ten is not face card",
			card:     Card{Suit: Hearts, Rank: Ten},
			expected: false,
		},
		{
			name:     "Two is not face card",
			card:     Card{Suit: Hearts, Rank: Two},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.card.IsFaceCard()
			if got != test.expected {
				t.Errorf("Card.IsFaceCard() = %v, want %v", got, test.expected)
			}
		})
	}
}

// TestIsAce tests ace detection
func TestIsAce(t *testing.T) {
	tests := []struct {
		name     string
		card     Card
		expected bool
	}{
		{
			name:     "Ace is ace",
			card:     Card{Suit: Hearts, Rank: Ace},
			expected: true,
		},
		{
			name:     "King is not ace",
			card:     Card{Suit: Hearts, Rank: King},
			expected: false,
		},
		{
			name:     "Two is not ace",
			card:     Card{Suit: Hearts, Rank: Two},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.card.IsAce()
			if got != test.expected {
				t.Errorf("Card.IsAce() = %v, want %v", got, test.expected)
			}
		})
	}
}

// TestCardString tests string representation
func TestCardString(t *testing.T) {
	tests := []struct {
		name     string
		card     Card
		expected string
	}{
		{
			name:     "Ace of Hearts",
			card:     Card{Suit: Hearts, Rank: Ace},
			expected: "Ace of Hearts",
		},
		{
			name:     "King of Spades",
			card:     Card{Suit: Spades, Rank: King},
			expected: "King of Spades",
		},
		{
			name:     "Ten of Diamonds",
			card:     Card{Suit: Diamonds, Rank: Ten},
			expected: "Ten of Diamonds",
		},
		{
			name:     "Two of Clubs",
			card:     Card{Suit: Clubs, Rank: Two},
			expected: "Two of Clubs",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.card.String()
			if got != test.expected {
				t.Errorf("Card.String() = %v, want %v", got, test.expected)
			}
		})
	}
}

// TestCardShortString tests short string representation
func TestCardShortString(t *testing.T) {
	tests := []struct {
		name     string
		card     Card
		expected string
	}{
		{
			name:     "Ace of Hearts",
			card:     Card{Suit: Hearts, Rank: Ace},
			expected: "A♥",
		},
		{
			name:     "King of Spades",
			card:     Card{Suit: Spades, Rank: King},
			expected: "K♠",
		},
		{
			name:     "Ten of Diamonds",
			card:     Card{Suit: Diamonds, Rank: Ten},
			expected: "10♦",
		},
		{
			name:     "Two of Clubs",
			card:     Card{Suit: Clubs, Rank: Two},
			expected: "2♣",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.card.ShortString()
			if got != test.expected {
				t.Errorf("Card.ShortString() = %v, want %v", got, test.expected)
			}
		})
	}
}
