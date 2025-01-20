# Memory Game in Go

This is a simple memory game implemented in Go. The objective of the game is to find pairs of matching cards on a 4x4 board. The game is played in a terminal and allows players to reveal cards, attempt to find pairs, and track their performance in terms of moves and time.

## Features

- **4x4 Board**: The game presents a 4x4 board with 16 cards, where each card is represented by a letter (A, B, C, D, E, F, G, H).
- **Card Revelation**: Players can choose two cards at a time to reveal. If the cards match, they remain revealed; otherwise, they are hidden again.
- **Move Count**: The game counts how many moves the player has made to find all pairs.
- **Scoring**: The score is calculated based on the time taken and the number of moves. The player starts with 10 points and loses points for each move and over time.
- **Board Rotation**: After every 5 moves, the player has the option to rotate the board, adding an element of challenge to the game.

## Requirements

- Go (version 1.16 or higher)

## How to Run

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your_username/memory-game.git
   cd memory-game
