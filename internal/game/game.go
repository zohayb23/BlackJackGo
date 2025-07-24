// Package game handles the core BlackJack game logic
package game

import (
	"blackjack/internal/deck"
	"blackjack/internal/player"
	"fmt"
)

// GameState represents the current state of the game
type GameState int

const (
	// Game states
	WaitingToStart GameState = iota // Waiting for the game to start
	PlayerTurn                      // Player's turn to act
	DealerTurn                      // Dealer's turn to act
	RoundOver                       // Round is complete
)

// Game represents a BlackJack game session
type Game struct {
	player *player.Player // The player
	dealer *player.Player // The dealer
	deck   *deck.Deck     // The game's deck
	state  GameState      // Current game state
}

// NewGame creates a new BlackJack game
func NewGame(playerName string) *Game {
	game := &Game{
		player: player.NewPlayer(playerName),
		dealer: player.NewPlayer("Dealer"),
		deck:   deck.NewDeck(),
		state:  WaitingToStart,
	}

	// Shuffle the deck
	game.deck.Shuffle()

	return game
}

// StartRound begins a new round of BlackJack
func (g *Game) StartRound() error {
	// Reset hands
	g.player.ClearHand()
	g.dealer.ClearHand()

	// Check if we need a new deck (less than 20 cards remaining)
	if g.deck.RemainingCards() < 20 {
		g.deck = deck.NewDeck()
		g.deck.Shuffle()
	}

	// Deal initial cards
	// First card to player and dealer
	card, err := g.deck.DrawCard()
	if err != nil {
		return fmt.Errorf("failed to deal card: %v", err)
	}
	g.player.AddCard(card)

	card, err = g.deck.DrawCard()
	if err != nil {
		return fmt.Errorf("failed to deal card to dealer: %v", err)
	}
	g.dealer.AddCard(card)

	// Second card to player and dealer
	card, err = g.deck.DrawCard()
	if err != nil {
		return fmt.Errorf("failed to deal card: %v", err)
	}
	g.player.AddCard(card)

	card, err = g.deck.DrawCard()
	if err != nil {
		return fmt.Errorf("failed to deal card to dealer: %v", err)
	}
	g.dealer.AddCard(card)

	// Check for player BlackJack
	if g.player.HasBlackjack() {
		g.player.Stand()
		g.state = DealerTurn
	} else {
		g.state = PlayerTurn
	}

	return nil
}

// GetDealerVisibleCard returns the dealer's face-up card
func (g *Game) GetDealerVisibleCard() (deck.Card, error) {
	if len(g.dealer.Hand) == 0 {
		return deck.Card{}, fmt.Errorf("dealer has no cards")
	}
	return g.dealer.Hand[0], nil
}

// PlayerHit handles the player's request to hit (take another card)
func (g *Game) PlayerHit() error {
	if g.state != PlayerTurn {
		return fmt.Errorf("cannot hit: not player's turn")
	}

	card, err := g.deck.DrawCard()
	if err != nil {
		return fmt.Errorf("failed to draw card: %v", err)
	}

	g.player.AddCard(card)

	// Check if player busted
	if g.player.State == player.Busted {
		g.state = RoundOver
	}

	return nil
}

// PlayerStand handles the player's request to stand (keep current hand)
func (g *Game) PlayerStand() error {
	if g.state != PlayerTurn {
		return fmt.Errorf("cannot stand: not player's turn")
	}

	g.player.Stand()
	g.state = DealerTurn
	return nil
}

// DealerPlay handles the dealer's turn
func (g *Game) DealerPlay() error {
	if g.state != DealerTurn {
		return fmt.Errorf("not dealer's turn")
	}

	// Dealer must hit on 16 and below, stand on 17 and above
	for g.dealer.GetHandValue() < 17 {
		card, err := g.deck.DrawCard()
		if err != nil {
			return fmt.Errorf("failed to draw card: %v", err)
		}
		g.dealer.AddCard(card)
	}

	g.state = RoundOver
	return nil
}

// GetResult returns the game result from the player's perspective
func (g *Game) GetResult() string {
	if g.state != RoundOver {
		return "Round not over"
	}

	playerValue := g.player.GetHandValue()
	dealerValue := g.dealer.GetHandValue()

	switch {
	case g.player.State == player.Busted:
		return "Player busted! Dealer wins!"
	case g.dealer.State == player.Busted:
		return "Dealer busted! Player wins!"
	case g.player.HasBlackjack() && !g.dealer.HasBlackjack():
		return "BlackJack! Player wins!"
	case !g.player.HasBlackjack() && g.dealer.HasBlackjack():
		return "Dealer has BlackJack! Dealer wins!"
	case playerValue > dealerValue:
		return "Player wins!"
	case dealerValue > playerValue:
		return "Dealer wins!"
	default:
		return "Push - It's a tie!"
	}
}

// GetState returns the current game state
func (g *Game) GetState() GameState {
	return g.state
}

// String returns a string representation of the game state
func (g *Game) String() string {
	dealerInfo := "Dealer: "
	if g.state == RoundOver {
		dealerInfo += g.dealer.String()
	} else {
		// Show only the first card during play
		if len(g.dealer.Hand) > 0 {
			dealerInfo += fmt.Sprintf("Shows %s (Hidden card)", g.dealer.Hand[0].String())
		}
	}

	return fmt.Sprintf("Game State: %v\n%s\nPlayer: %s",
		g.state, dealerInfo, g.player.String())
}
