package main

import (
  "fmt"
  "math/rand"
  "time"
)

const (
  boardSize = 4
)

var (
  board       [boardSize][boardSize]string
  revealed    [boardSize][boardSize]bool
  cardValues  []string
  moves       int
  startTime   time.Time
  score       int
)

func initBoard() {
  cardValues = []string{"A", "A", "B", "B", "C", "C", "D", "D", "E", "E", "F", "F", "G", "G", "H", "H"}
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(cardValues), func(i, j int) {
    cardValues[i], cardValues[j] = cardValues[j], cardValues[i]
  })

  for i := 0; i < boardSize; i++ {
    for j := 0; j < boardSize; j++ {
      board[i][j] = cardValues[i*boardSize+j]
    }
  }
}

func displayBoard() {
  fmt.Println("Tabuleiro:")
  for i := 0; i < boardSize; i++ {
    for j := 0; j < boardSize; j++ {
      if revealed[i][j] {
        fmt.Print(board[i][j], " ")
      } else {
        fmt.Print("X ")
      }
    }
    fmt.Println()
  }
  fmt.Println()
}

func allRevealed() bool {
  for i := 0; i < boardSize; i++ {
    for j := 0; j < boardSize; j++ {
      if !revealed[i][j] {
        return false
      }
    }
  }
  return true
}

func rotateBoard() {
  var newBoard [boardSize][boardSize]string
  for i := 0; i < boardSize; i++ {
    for j := 0; j < boardSize; j++ {
      newBoard[j][boardSize-1-i] = board[i][j]
    }
  }
  board = newBoard
}

func calculateScore() int {
  elapsed := time.Since(startTime).Seconds()
  score = 10 - moves - int(elapsed)
  if score < 0 {
    score = 0
  }
  return score
}

func playRound() {
  var x1, y1, x2, y2 int
  startTime = time.Now()
  for {
    displayBoard()
    fmt.Printf("Escolha a primeira carta (linha coluna) ou digite -1 para desistir: ")
    fmt.Scanf("%d %d", &x1, &y1)
    if x1 == -1 {
      fmt.Printf("Você desistiu! Tempo: %.2f segundos, Pontuação: %d\n", time.Since(startTime).Seconds(), calculateScore())
      break
    }
    x1--
    y1--

    if x1 < 0 || x1 >= boardSize || y1 < 0 || y1 >= boardSize || revealed[x1][y1] {
      fmt.Println("Movimento inválido! Tente novamente.")
      continue
    }

    revealed[x1][y1] = true
    displayBoard()

    fmt.Printf("Escolha a segunda carta (linha coluna): ")
    fmt.Scanf("%d %d", &x2, &y2)
    x2--
    y2--

    if x2 < 0 || x2 >= boardSize || y2 < 0 || y2 >= boardSize || revealed[x2][y2] || (x1 == x2 && y1 == y2) {
      fmt.Println("Movimento inválido! Tente novamente.")
      revealed[x1][y1] = false
      continue
    }

    revealed[x2][y2] = true
    displayBoard()
    moves++

    if board[x1][y1] != board[x2][y2] {
      fmt.Println("Não é um par! Tente novamente.")
      revealed[x1][y1] = false
      revealed[x2][y2] = false
    } else {
      fmt.Println("Par encontrado!")
    }

    if allRevealed() {
      fmt.Printf("Você ganhou em %d movimentos! Tempo: %.2f segundos, Pontuação: %d\n", moves, time.Since(startTime).Seconds(), calculateScore())
      break
    }

    if moves%5 == 0 {
      var rotate int
      fmt.Printf("Deseja girar o tabuleiro? (1 para sim, 0 para não): ")
      fmt.Scanf("%d", &rotate)
      if rotate == 1 {
        rotateBoard()
      }
    }
  }
}

func main() {
  initBoard()
  playRound()
}
