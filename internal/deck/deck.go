package deck

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Deck represents a collection of playing cards
type Deck struct {
	cards []Card // Using a slice to store cards
}

// NewDeck creates a new standard deck of 52 playing cards
func NewDeck() *Deck {
	// Create an empty deck
	d := &Deck{
		cards: make([]Card, 0, 52), // Initialize with 0 length but capacity of 52
	}

	// Define all suits and ranks
	suits := []Suit{Hearts, Diamonds, Clubs, Spades}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	// Add all combinations of suits and ranks
	for _, suit := range suits {
		for _, rank := range ranks {
			// Create a new card and add it to the deck
			card, _ := NewCard(suit, rank) // We can ignore the error here because we know these are valid
			d.cards = append(d.cards, card)
		}
	}

	return d
}

// Shuffle randomizes the order of cards in the deck
// Makes sure that the deck is shuffled before drawing cards
func (d *Deck) Shuffle() {
	// Create a new random source with current time as seed to ensure randomness
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Use the Fisher-Yates shuffle algorithm to randomize the order of cards
	for i := len(d.cards) - 1; i > 0; i-- {
		// Generate a random index between 0 and i
		j := r.Intn(i + 1)
		// Swap cards[i] with cards[j]
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// DrawCard removes and returns the top card from the deck 
// In Go, we can return multiple values - here we return both the card and an error (if the deck is empty)
func (d *Deck) DrawCard() (Card, error) {
	// Check if deck is empty 
	if len(d.cards) == 0 {
		return Card{}, fmt.Errorf("cannot draw from empty deck")
	}

	// Get the top card 
	card := d.cards[0]
	// Remove it from the deck (slice from index 1 onwards)
	d.cards = d.cards[1:]
	return card, nil
}

// RemainingCards returns the number of cards left in the deck
func (d *Deck) RemainingCards() int {
	return len(d.cards)
}

// String returns a string representation of the deck
func (d *Deck) String() string {
	// Create a string builder for efficient string concatenation
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Deck with %d cards:\n", len(d.cards)))

	// Loop through cards and add their string representations
	for i, card := range d.cards {
		sb.WriteString(fmt.Sprintf("%d: %s\n", i+1, card.String()))
	}

	return sb.String()
}
