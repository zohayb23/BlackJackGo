package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"blackjack/internal/game"
	"blackjack/internal/rules"
)

// clearScreen clears the terminal screen
func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen
}

// getPlayerName prompts for and returns the player's name
func getPlayerName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your name: ")
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}

// getPlayerInput reads and returns the player's command
func getPlayerInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter command (h/hit, s/stand, r/rules, q/quit): ")
	input, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(input))
}

// displayGameState shows the current state of the game
func displayGameState(g *game.Game) {
	clearScreen()
	fmt.Println("\n=== BLACKJACK ===")
	fmt.Println(g.String())
}

// playRound plays a single round of BlackJack
func playRound(g *game.Game) bool {
	err := g.StartRound()
	if err != nil {
		fmt.Printf("Error starting round: %v\n", err)
		return false
	}

	// Main game loop
	for {
		displayGameState(g)

		// Check if player's turn is over
		if g.GetState() == game.RoundOver {
			fmt.Println("\n" + g.GetResult())
			return true
		}

		if g.GetState() == game.DealerTurn {
			fmt.Println("\nDealer's turn...")
			err := g.DealerPlay()
			if err != nil {
				fmt.Printf("Error during dealer play: %v\n", err)
				return false
			}
			continue
		}

		// Get player command
		cmd := getPlayerInput()
		switch cmd {
		case "h", "hit":
			err := g.PlayerHit()
			if err != nil {
				fmt.Printf("Error hitting: %v\n", err)
			}

		case "s", "stand":
			err := g.PlayerStand()
			if err != nil {
				fmt.Printf("Error standing: %v\n", err)
			}

		case "r", "rules":
			clearScreen()
			fmt.Println(rules.DisplayAllRules())
			fmt.Println(rules.DisplayHelp())
			fmt.Println("\nPress Enter to continue...")
			bufio.NewReader(os.Stdin).ReadString('\n')

		case "q", "quit":
			return false

		default:
			fmt.Println("Invalid command. Try again.")
		}
	}
}

func main() {
	clearScreen()
	fmt.Println(rules.DisplayAllRules())
	fmt.Println(rules.DisplayHelp())
	fmt.Println("\nPress Enter to start...")
	bufio.NewReader(os.Stdin).ReadString('\n')

	name := getPlayerName()
	g := game.NewGame(name)

	// Main game loop
	for {
		if !playRound(g) {
			break
		}

		// Ask to play another round
		fmt.Print("\nPlay another round? (y/n): ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.ToLower(strings.TrimSpace(input)) != "y" {
			break
		}
	}

	fmt.Println("\nThanks for playing!")
}
