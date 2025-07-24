package rules

import (
	"strings"
	"testing"
)

// TestGetGameRules tests the game rules content
func TestGetGameRules(t *testing.T) {
	rules := GetGameRules()

	// Check that we have all required sections
	expectedTitles := []string{
		"Game Objective",
		"Card Values",
		"Game Flow",
		"Winning Conditions",
	}

	if len(rules) != len(expectedTitles) {
		t.Errorf("Expected %d rule sections, got %d", len(expectedTitles), len(rules))
	}

	// Check each section has required content
	for i, section := range rules {
		if section.Title != expectedTitles[i] {
			t.Errorf("Expected section title %q, got %q", expectedTitles[i], section.Title)
		}
		if len(section.Content) == 0 {
			t.Errorf("Section %q has no content", section.Title)
		}
	}
}

// TestGetCommandHelp tests the command help content
func TestGetCommandHelp(t *testing.T) {
	help := GetCommandHelp()

	if len(help) == 0 {
		t.Error("Expected command help sections, got none")
	}

	// Check that all important commands are mentioned
	commandSection := help[0]
	requiredCommands := []string{
		"hit",
		"stand",
		"rules",
		"quit",
	}

	for _, cmd := range requiredCommands {
		if !strings.Contains(strings.ToLower(commandSection.Content), cmd) {
			t.Errorf("Command help missing %q command", cmd)
		}
	}
}

// TestDisplaySection tests section formatting
func TestDisplaySection(t *testing.T) {
	section := Section{
		Title:   "Test Title",
		Content: "Test Content",
	}

	result := DisplaySection(section)

	expectedParts := []string{
		"===",          // Section delimiter
		"Test Title",   // Title
		"Test Content", // Content
	}

	for _, part := range expectedParts {
		if !strings.Contains(result, part) {
			t.Errorf("Formatted section missing %q", part)
		}
	}
}

// TestDisplayAllRules tests complete rules display
func TestDisplayAllRules(t *testing.T) {
	rules := DisplayAllRules()

	// Check main header
	if !strings.Contains(rules, "BLACKJACK RULES") {
		t.Error("Rules display missing main header")
	}

	// Check that all sections are included
	for _, section := range GetGameRules() {
		if !strings.Contains(rules, section.Title) {
			t.Errorf("Rules display missing section %q", section.Title)
		}
		if !strings.Contains(rules, section.Content) {
			t.Errorf("Rules display missing content for section %q", section.Title)
		}
	}
}

// TestDisplayHelp tests help text display
func TestDisplayHelp(t *testing.T) {
	help := DisplayHelp()

	// Check main header
	if !strings.Contains(help, "GAME HELP") {
		t.Error("Help display missing main header")
	}

	// Check that command section is included
	commandHelp := GetCommandHelp()[0]
	if !strings.Contains(help, commandHelp.Title) {
		t.Error("Help display missing command section")
	}
	if !strings.Contains(help, commandHelp.Content) {
		t.Error("Help display missing command content")
	}
}
