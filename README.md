# BlackJack in Go

A command-line implementation of the BlackJack card game using Go. This project serves as a learning exercise for Go programming concepts and best practices.

## Project Overview

This BlackJack implementation showcases:

- Object-oriented programming in Go
- Clean code architecture
- Test-driven development
- Command-line interface design
- Professional development practices

## Current Progress

### Completed Features ✓

- Basic project structure and Git setup
- Card implementation with:
  - Value calculation
  - Suit and rank representation
  - String formatting
  - Validation
  - 100% test coverage
- Deck implementation with:
  - Standard 52-card deck creation
  - Shuffling functionality
  - Card drawing
  - Remaining cards tracking
  - 100% test coverage
- Player implementation with:
  - Hand management
  - State tracking (Playing, Standing, Busted, BlackJack)
  - Card value calculation with Ace handling
  - 97% test coverage
- Game logic with:
  - Single player vs dealer
  - State management
  - Dealer AI (hit on 16, stand on 17)
  - Win condition checking
  - 100% test coverage
- Game rules and documentation:
  - Comprehensive rules text
  - Command help
  - Formatted display
  - 100% test coverage
- Command-line interface:
  - Interactive gameplay
  - Clear screen management
  - Hidden dealer card
  - Game state display
  - Command processing
  - Round management

### Project Structure

```
blackjack/
├── cmd/            # Command-line application entry point
│   └── main.go     # Main game interface
├── internal/       # Private application code
│   ├── deck/      # Card and deck implementations
│   │   ├── card.go    # Card struct and methods
│   │   └── deck.go    # Deck struct and methods
│   ├── game/      # Game logic
│   │   └── game.go    # Game struct and methods
│   ├── player/    # Player implementation
│   │   └── player.go  # Player struct and methods
│   └── rules/     # Game rules and help text
│       └── rules.go   # Rules content and formatting
├── docs/          # Documentation
└── pkg/           # Public packages (if any)
```

## Technical Highlights

### Go Concepts Implemented

- Custom types and structs
- Methods and interfaces
- Error handling
- Slices and maps
- Package organization
- State management
- String formatting
- User input handling
- ANSI terminal control

### Testing Coverage

- deck/card.go: 100% coverage
- deck/deck.go: 100% coverage
- player/player.go: 97% coverage
- game/game.go: 100% coverage
- rules/rules.go: 100% coverage

### Development Practices

- Clear commit messages
- Incremental development
- Test-driven development
- Comprehensive documentation
- Code organization
- Error handling

## Getting Started

### Prerequisites

- Go 1.24.5 or later

### Installation

```bash
# Clone the repository
git clone https://github.com/zohayb23/BlackJackGo.git

# Navigate to project directory
cd BlackJackGo

# Run the game
go run cmd/main.go
```

### How to Play

1. Start the game
2. Enter your name
3. Use the following commands:
   - `h` or `hit` - Take another card
   - `s` or `stand` - Keep your current hand
   - `r` or `rules` - Display game rules
   - `q` or `quit` - Exit the game

## Documentation

- See [docs/LEARNING.txt](docs/LEARNING.txt) for detailed Go concepts covered
- Game rules are available in-game via the 'r' command

## Contributing

This is a learning project, but suggestions and feedback are welcome!

## License

This project is open source and available under the MIT License.
