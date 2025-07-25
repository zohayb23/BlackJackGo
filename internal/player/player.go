// Package player handles the player-related functionality in the game
package player

import (
	"blackjack/internal/deck"
	"fmt"
)

// PlayerState represents the current state of a player in the game
type PlayerState int

const (
	// Different states a player can be in during the game
	//iota is a special keyword that automatically assigns sequential integer values to constants starting from 0
	Playing   PlayerState = iota // Player is still making decisions
	Standing                     // Player has chosen to stand
	Busted                       // Player's hand value exceeded 21
	BlackJack                    // Player has a BlackJack (21 with 2 cards)
)

// Player represents a player in the game
type Player struct {
	Name  string      // Player's name
	Hand  []deck.Card // Cards in player's hand
	State PlayerState // Current state of the player
}

// NewPlayer creates a new player with the given name
func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Hand:  make([]deck.Card, 0), // Initialize empty hand
		State: Playing,              // Start in Playing state
	}
}

// AddCard adds a card to the player's hand
func (p *Player) AddCard(card deck.Card) {
	p.Hand = append(p.Hand, card)

	// After adding a card, check if player has busted
	if p.GetHandValue() > 21 {
		p.State = Busted
	} else if len(p.Hand) == 2 && p.GetHandValue() == 21 {
		p.State = BlackJack
	}
}

// GetHandValue calculates the total value of the player's hand
func (p *Player) GetHandValue() int {
	total := 0
	aceCount := 0

	// First pass: count all cards, treating Aces as 11
	for _, card := range p.Hand {
		if card.IsAce() {
			aceCount++
			total += 11
		} else {
			total += card.Value()
		}
	}

	// Second pass: convert Aces from 11 to 1 if we're over 21
	for aceCount > 0 && total > 21 {
		total -= 10 // Convert one Ace from 11 to 1
		aceCount--
	}

	return total
}

// Stand changes the player's state to Standing
func (p *Player) Stand() {
	p.State = Standing
}

// ClearHand removes all cards from the player's hand
func (p *Player) ClearHand() {
	p.Hand = p.Hand[:0] // Clear slice while preserving capacity
	p.State = Playing   // Reset state to Playing
}

// HasBlackjack checks if the player has a natural blackjack (21 with 2 cards)
func (p *Player) HasBlackjack() bool {
	return len(p.Hand) == 2 && p.GetHandValue() == 21
}

// String returns a string representation of the player's current state
func (p *Player) String() string {
	handStr := ""
	for i, card := range p.Hand {
		if i > 0 {
			handStr += ", "
		}
		handStr += card.String()
	}

	return fmt.Sprintf("Player: %s\nHand: %s\nValue: %d\nState: %v",
		p.Name, handStr, p.GetHandValue(), p.State)
}

// String returns a string representation of the PlayerState
func (s PlayerState) String() string {
	switch s {
	case Playing:
		return "Playing"
	case Standing:
		return "Standing"
	case Busted:
		return "Busted"
	case BlackJack:
		return "BlackJack"
	default:
		return "Unknown" // Default case if the state is not one of the above	
	}
}
