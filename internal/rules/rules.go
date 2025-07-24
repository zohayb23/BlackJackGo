// Package rules provides game rules and help text for BlackJack
package rules

import "fmt"

// Section represents a section of rules or help text
type Section struct {
	Title   string
	Content string
}

// GetGameRules returns all game rules sections
func GetGameRules() []Section {
	return []Section{
		{
			Title: "Game Objective",
			Content: `The goal is to beat the dealer by:
• Getting a hand value closer to 21 than the dealer
• Having the dealer go over 21 (bust)
• Getting a BlackJack (Ace + 10-value card) when dealer doesn't`,
		},
		{
			Title: "Card Values",
			Content: `• Ace: 11 or 1 (automatically adjusted to prevent busting)
• Face Cards (Jack, Queen, King): 10
• Number Cards: Their face value (2-10)`,
		},
		{
			Title: "Game Flow",
			Content: `1. You and the dealer each get two cards
2. One of dealer's cards remains hidden until your turn ends
3. You can repeatedly choose to:
   • Hit - Take another card
   • Stand - Keep your current hand
4. If you go over 21, you bust and lose
5. If you stand, the dealer reveals their hidden card
6. Dealer must hit on 16 or below, and stand on 17 or above`,
		},
		{
			Title: "Winning Conditions",
			Content: `You win if:
• You get a BlackJack (Ace + 10-value card)
• Your final hand is closer to 21 than the dealer
• Dealer busts (goes over 21)

You lose if:
• You bust (go over 21)
• Dealer's hand is closer to 21 than yours
• Dealer gets BlackJack when you don't

If both hands are equal, it's a tie (Push)`,
		},
	}
}

// GetCommandHelp returns help text for game commands
func GetCommandHelp() []Section {
	return []Section{
		{
			Title: "Game Commands",
			Content: `Available commands during play:
• h or hit   - Take another card
• s or stand - Keep your current hand
• r or rules - Display game rules
• q or quit  - Exit the game`,
		},
	}
}

// DisplaySection formats and prints a single section
func DisplaySection(section Section) string {
	return fmt.Sprintf("\n=== %s ===\n%s\n", section.Title, section.Content)
}

// DisplayAllRules formats and returns the complete rules text
func DisplayAllRules() string {
	var result string
	result += "\n=== BLACKJACK RULES ===\n"

	for _, section := range GetGameRules() {
		result += DisplaySection(section)
	}

	return result
}

// DisplayHelp formats and returns the command help text
func DisplayHelp() string {
	var result string
	result += "\n=== GAME HELP ===\n"

	for _, section := range GetCommandHelp() {
		result += DisplaySection(section)
	}

	return result
}
