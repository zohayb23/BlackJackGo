// Package deck handles the playing card functionality
package deck

import "fmt"

// Suit represents the suit of a playing card (Hearts, Diamonds, etc.)
type Suit string

// Rank represents the rank of a playing card (Ace, Two, Three, etc.)
type Rank string

// These constants define the four possible suits in a deck
const (
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
	Spades   Suit = "Spades"
)

// These constants define all possible ranks in a deck
const (
	Ace   Rank = "Ace"
	Two   Rank = "Two"
	Three Rank = "Three"
	Four  Rank = "Four"
	Five  Rank = "Five"
	Six   Rank = "Six"
	Seven Rank = "Seven"
	Eight Rank = "Eight"
	Nine  Rank = "Nine"
	Ten   Rank = "Ten"
	Jack  Rank = "Jack"
	Queen Rank = "Queen"
	King  Rank = "King"
)

// suitSymbols maps suits to their Unicode symbols
var suitSymbols = map[Suit]string{
	Hearts:   "♥",
	Diamonds: "♦",
	Clubs:    "♣",
	Spades:   "♠",
}

// rankSymbols maps ranks to their short representations
var rankSymbols = map[Rank]string{
	Ace:   "A",
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
}

// Card represents a playing card with a suit and rank
type Card struct {
	Suit Suit
	Rank Rank
}

// NewCard creates a new card and validates the suit and rank
// In Go, this is called a constructor function (though it's just a regular function)
func NewCard(suit Suit, rank Rank) (Card, error) {
	// Check if the suit is valid
	if _, ok := suitSymbols[suit]; !ok {
		return Card{}, fmt.Errorf("invalid suit: %s", suit)
	}

	// Check if the rank is valid
	if _, ok := rankSymbols[rank]; !ok {
		return Card{}, fmt.Errorf("invalid rank: %s", rank)
	}

	return Card{Suit: suit, Rank: rank}, nil
}

// Value returns the numerical value of the card in BlackJack
func (c Card) Value() int {
	switch c.Rank {
	case Ace:
		return 11 // In BlackJack, Ace can be 1 or 11 (we'll handle this logic in the game)
	case Ten, Jack, Queen, King:
		return 10
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	default:
		return 0
	}
}

// IsFaceCard returns true if the card is a Jack, Queen, or King
func (c Card) IsFaceCard() bool {
	return c.Rank == Jack || c.Rank == Queen || c.Rank == King
}

// IsAce returns true if the card is an Ace
func (c Card) IsAce() bool {
	return c.Rank == Ace
}

// String returns a string representation of the card (e.g., "Ace of Hearts")
func (c Card) String() string {
	return string(c.Rank) + " of " + string(c.Suit)
}

// ShortString returns a short representation of the card (e.g., "A♥")
func (c Card) ShortString() string {
	return rankSymbols[c.Rank] + suitSymbols[c.Suit]
}
