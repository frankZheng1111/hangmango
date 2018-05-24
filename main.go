package main

import (
  "fmt"
  game "hangmango/models"
)

func main() {
  hangman := game.StartNewGame("test")
  fmt.Println(hangman)
}
