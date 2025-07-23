package deck

import "testing"

// TestCardValue tests the Value() method of Card
func TestCardValue(t *testing.T) {
	// In Go, we can create a slice of test cases using struct
	tests := []struct {
		name     string    // name of the test case
		card     Card      // input card
		expected int       // expected value
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

	// Loop through all test cases
	for _, test := range tests {
		// t.Run creates a subtest with the given name
		t.Run(test.name, func(t *testing.T) {
			// Get the actual value
			got := test.card.Value()
			// Compare expected and actual values
			if got != test.expected {
				t.Errorf("Card.Value() = %v, want %v", got, test.expected)
			}
		})
	}
}

// TestCardString tests the String() method of Card
func TestCardString(t *testing.T) {
	// Create a test card
	card := Card{
		Suit: Hearts,
		Rank: Ace,
	}

	// Expected string representation
	expected := "Ace of Hearts"

	// Get the actual string
	got := card.String()

	// Compare expected and actual strings
	if got != expected {
		t.Errorf("Card.String() = %v, want %v", got, expected)
	}
} 