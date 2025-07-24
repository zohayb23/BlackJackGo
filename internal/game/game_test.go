package game

import (
	"blackjack/internal/deck"
	"strings"
	"testing"
)

// TestNewGame tests game creation
func TestNewGame(t *testing.T) {
	game := NewGame("Test Player")

	if game.state != WaitingToStart {
		t.Errorf("Expected initial state WaitingToStart, got %v", game.state)
	}
	if game.player == nil {
		t.Error("Expected player to be initialized")
	}
	if game.dealer == nil {
		t.Error("Expected dealer to be initialized")
	}
	if game.deck == nil {
		t.Error("Expected deck to be initialized")
	}
}

// TestStartRound tests starting a new round
func TestStartRound(t *testing.T) {
	game := NewGame("Test Player")

	err := game.StartRound()
	if err != nil {
		t.Errorf("Unexpected error starting round: %v", err)
	}

	// Check initial deal
	if len(game.player.Hand) != 2 {
		t.Error("Expected player to have 2 cards")
	}
	if len(game.dealer.Hand) != 2 {
		t.Error("Expected dealer to have 2 cards")
	}

	// Test state after dealing
	if game.state != PlayerTurn && game.state != DealerTurn {
		t.Error("Expected state to be PlayerTurn or DealerTurn (in case of player BlackJack)")
	}
}

// TestPlayerActions tests player hit and stand actions
func TestPlayerActions(t *testing.T) {
	t.Run("Player Hit", func(t *testing.T) {
		game := NewGame("Test Player")
		game.StartRound()

		initialCards := len(game.player.Hand)
		err := game.PlayerHit()

		if err != nil {
			t.Errorf("Unexpected error on hit: %v", err)
		}
		if len(game.player.Hand) != initialCards+1 {
			t.Error("Expected player to receive one card")
		}
	})

	t.Run("Player Stand", func(t *testing.T) {
		game := NewGame("Test Player")
		game.StartRound()

		err := game.PlayerStand()
		if err != nil {
			t.Errorf("Unexpected error on stand: %v", err)
		}
		if game.state != DealerTurn {
			t.Error("Expected state to change to DealerTurn")
		}
	})

	t.Run("Cannot Hit After Stand", func(t *testing.T) {
		game := NewGame("Test Player")
		game.StartRound()
		game.PlayerStand()

		err := game.PlayerHit()
		if err == nil {
			t.Error("Expected error when hitting after stand")
		}
	})
}

// TestDealerPlay tests dealer's turn
func TestDealerPlay(t *testing.T) {
	game := NewGame("Test Player")
	game.StartRound()
	game.PlayerStand()

	err := game.DealerPlay()
	if err != nil {
		t.Errorf("Unexpected error during dealer play: %v", err)
	}

	// Verify dealer follows rules (must have 17 or more)
	if game.dealer.GetHandValue() < 17 {
		t.Error("Dealer should hit until at least 17")
	}

	if game.state != RoundOver {
		t.Error("Game should be over after dealer plays")
	}
}

// TestGetResult tests game result determination
func TestGetResult(t *testing.T) {
	tests := []struct {
		name           string
		setupGame      func(*Game)
		expectedResult string
	}{
		{
			name: "Player Busts",
			setupGame: func(g *Game) {
				// Give player cards that will bust
				g.player.AddCard(mustCreateCard(t, deck.Hearts, deck.King))
				g.player.AddCard(mustCreateCard(t, deck.Spades, deck.Queen))
				g.player.AddCard(mustCreateCard(t, deck.Diamonds, deck.Jack))
				g.state = RoundOver
			},
			expectedResult: "Player busted",
		},
		{
			name: "Player BlackJack",
			setupGame: func(g *Game) {
				g.player.AddCard(mustCreateCard(t, deck.Hearts, deck.Ace))
				g.player.AddCard(mustCreateCard(t, deck.Spades, deck.King))
				g.dealer.AddCard(mustCreateCard(t, deck.Hearts, deck.Ten))
				g.dealer.AddCard(mustCreateCard(t, deck.Spades, deck.Nine))
				g.state = RoundOver
			},
			expectedResult: "BlackJack",
		},
		{
			name: "Push",
			setupGame: func(g *Game) {
				g.player.AddCard(mustCreateCard(t, deck.Hearts, deck.Ten))
				g.player.AddCard(mustCreateCard(t, deck.Spades, deck.Nine))
				g.dealer.AddCard(mustCreateCard(t, deck.Diamonds, deck.Ten))
				g.dealer.AddCard(mustCreateCard(t, deck.Clubs, deck.Nine))
				g.state = RoundOver
			},
			expectedResult: "Push",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := NewGame("Test Player")
			test.setupGame(game)

			result := game.GetResult()
			if !strings.Contains(result, test.expectedResult) {
				t.Errorf("Expected result containing %q, got %q", test.expectedResult, result)
			}
		})
	}
}

// TestString tests game state string representation
func TestString(t *testing.T) {
	game := NewGame("Test Player")
	game.StartRound()

	// During play, dealer's second card should be hidden
	str := game.String()
	if !strings.Contains(str, "Hidden card") {
		t.Error("Should indicate dealer has a hidden card")
	}

	// After round is over, all cards should be visible
	game.state = RoundOver
	str = game.String()
	if strings.Contains(str, "Hidden card") {
		t.Error("Should show all dealer cards after round is over")
	}
	if !strings.Contains(str, "Dealer") || !strings.Contains(str, "Player") {
		t.Error("Should show both dealer and player information")
	}
}

// TestScoreTracking tests score tracking functionality
func TestScoreTracking(t *testing.T) {
	tests := []struct {
		name           string
		playerCards    []deck.Card
		dealerCards    []deck.Card
		expectedScore  Score
		expectedResult string
	}{
		{
			name: "Player wins with higher value",
			playerCards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ten),
				mustCreateCard(t, deck.Spades, deck.Nine),
			},
			dealerCards: []deck.Card{
				mustCreateCard(t, deck.Diamonds, deck.Ten),
				mustCreateCard(t, deck.Clubs, deck.Eight),
			},
			expectedScore:  Score{Wins: 1},
			expectedResult: "Player wins!",
		},
		{
			name: "Dealer wins with higher value",
			playerCards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ten),
				mustCreateCard(t, deck.Spades, deck.Seven),
			},
			dealerCards: []deck.Card{
				mustCreateCard(t, deck.Diamonds, deck.Ten),
				mustCreateCard(t, deck.Clubs, deck.Eight),
			},
			expectedScore:  Score{Losses: 1},
			expectedResult: "Dealer wins!",
		},
		{
			name: "Push with equal values",
			playerCards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ten),
				mustCreateCard(t, deck.Spades, deck.Eight),
			},
			dealerCards: []deck.Card{
				mustCreateCard(t, deck.Diamonds, deck.Ten),
				mustCreateCard(t, deck.Clubs, deck.Eight),
			},
			expectedScore:  Score{Pushes: 1},
			expectedResult: "Push! It's a tie!",
		},
		{
			name: "Player wins with BlackJack",
			playerCards: []deck.Card{
				mustCreateCard(t, deck.Hearts, deck.Ace),
				mustCreateCard(t, deck.Spades, deck.King),
			},
			dealerCards: []deck.Card{
				mustCreateCard(t, deck.Diamonds, deck.Ten),
				mustCreateCard(t, deck.Clubs, deck.Eight),
			},
			expectedScore:  Score{Wins: 1},
			expectedResult: "BlackJack! Player wins!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGame("TestPlayer")
			g.state = RoundOver

			// Set up player and dealer hands
			for _, card := range tt.playerCards {
				g.player.AddCard(card)
			}
			for _, card := range tt.dealerCards {
				g.dealer.AddCard(card)
			}

			// Get result and check score
			result := g.GetResult()
			if result != tt.expectedResult {
				t.Errorf("Expected result %q, got %q", tt.expectedResult, result)
			}

			score := g.GetScore()
			if score != tt.expectedScore {
				t.Errorf("Expected score %+v, got %+v", tt.expectedScore, score)
			}
		})
	}
}

func TestScoreDisplay(t *testing.T) {
	g := NewGame("TestPlayer")
	g.score = Score{Wins: 2, Losses: 1, Pushes: 1}

	output := g.String()
	if !strings.Contains(output, "Session Score - Wins: 2, Losses: 1, Pushes: 1") {
		t.Errorf("Score not displayed correctly in game state, got: %s", output)
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
