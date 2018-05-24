package models

import (
  "hangmango/config"
  "strings"
  _ "fmt"
)

type Hangman struct {
  ProtoWord string
  GuessedLetters map[string]bool
  WordLetters map[string]bool
  Hp int
}

func StartNewGame(customWord string) *Hangman{
  hangman := new(Hangman)
  hangman.ProtoWord = customWord
  hangman.WordLetters = make(map[string]bool)
  hangman.GuessedLetters = make(map[string]bool)
  hangman.Hp = config.GameSettingConfig.Hp
  for _, letter := range strings.Split(customWord, "") {
    hangman.WordLetters[letter] = false
  }
  return hangman
}

func (hangman *Hangman) isAlive() bool {
  return hangman.Hp > 0
}

func (hangman *Hangman) isWin() bool {
  for _, hasGuessed := range hangman.WordLetters {
    if !hasGuessed {
      return false
    }
  }
  return true
}
