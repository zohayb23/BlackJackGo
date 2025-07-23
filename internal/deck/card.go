// Package deck handles the playing card functionality
package deck

// Suit represents the suit of a playing card (Hearts, Diamonds, etc.)
// In Go, we can create our own types using the 'type' keyword
type Suit string

// Rank represents the rank of a playing card (Ace, Two, Three, etc.)
type Rank string

// These constants define the four possible suits in a deck
// In Go, constants are declared using the 'const' keyword
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

// Card represents a playing card with a suit and rank
// In Go, we create structs to combine different types of data
type Card struct {
	Suit Suit // The suit of the card (Hearts, Diamonds, etc.)
	Rank Rank // The rank of the card (Ace, Two, Three, etc.)
}

// Value returns the numerical value of the card in BlackJack
// In Go, we can add methods to our types using this syntax: func (receiverName ReceiverType) MethodName()
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
		return 0 // This should never happen with valid cards
	}
}

// String returns a string representation of the card (e.g., "Ace of Hearts")
// This is a special method in Go that's called when you print the card
func (c Card) String() string {
	return string(c.Rank) + " of " + string(c.Suit)
} 