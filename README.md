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
- Deck implementation with:
  - Standard 52-card deck creation
  - Shuffling functionality
  - Card drawing
  - Remaining cards tracking
- Comprehensive test coverage

### In Progress 🚧

- Player implementation
- Game logic
- Command-line interface

### Upcoming Features 📋

- Betting system
- Multiple players support
- Game statistics
- Save/load functionality

## Project Structure

```
blackjack/
├── cmd/            # Command-line application entry point
├── internal/       # Private application code
│   ├── deck/      # Card and deck implementations
│   ├── game/      # Game logic (upcoming)
│   └── player/    # Player implementation (upcoming)
├── pkg/           # Public packages (if any)
└── docs/          # Documentation
```

## Technical Highlights

### Go Concepts Implemented

- Custom types and structs
- Methods and interfaces
- Error handling
- Slices and maps
- Package organization
- Testing
- And more (see docs/LEARNING.txt)

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

# Run the tests
go test ./...
```

## Documentation

- See [docs/LEARNING.txt](docs/LEARNING.txt) for detailed Go concepts covered
- More documentation will be added as the project progresses

## Contributing

This is a learning project, but suggestions and feedback are welcome!

## License

This project is open source and available under the MIT License.
