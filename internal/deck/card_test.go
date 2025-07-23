package deck

import "testing"

// TestNewCard tests the NewCard constructor function
func TestNewCard(t *testing.T) {
	// Test cases for NewCard
	tests := []struct {
		name        string
		suit        Suit
		rank        Rank
		expectError bool
	}{
		{
			name:        "Valid card",
			suit:        Hearts,
			rank:        Ace,
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			card, err := NewCard(test.suit, test.rank)

			// Check if we got an error when we expected one
			if test.expectError && err == nil {
				t.Error("Expected an error but didn't get one")
			}

			// Check if we got an error when we didn't expect one
			if !test.expectError && err != nil {
				t.Errorf("Didn't expect an error but got: %v", err)
			}

			// For valid cards, check the values
			if err == nil {
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

// TestCardValue tests the Value() method of Card
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
			name:     "Number card value",
			card:     Card{Suit: Diamonds, Rank: Five},
			expected: 5,
		},
		{
			name:     "Face card value",
			card:     Card{Suit: Clubs, Rank: King},
			expected: 10,
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

// TestIsFaceCard tests the IsFaceCard() method
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
			name:     "Ace is not face card",
			card:     Card{Suit: Hearts, Rank: Ace},
			expected: false,
		},
		{
			name:     "Number is not face card",
			card:     Card{Suit: Hearts, Rank: Five},
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

// TestIsAce tests the IsAce() method
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

// TestCardString tests the String() method
func TestCardString(t *testing.T) {
	card := Card{Suit: Hearts, Rank: Ace}
	expected := "Ace of Hearts"
	got := card.String()
	if got != expected {
		t.Errorf("Card.String() = %v, want %v", got, expected)
	}
}

// TestCardShortString tests the ShortString() method
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
			name:     "Ten of Diamonds",
			card:     Card{Suit: Diamonds, Rank: Ten},
			expected: "10♦",
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
