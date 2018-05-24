package models

import (
	_ "fmt"
	"hangmango/config"
	"strings"
)

type Hangman struct {
	ProtoWord      string
	GuessedLetters map[string]bool
	WordLetters    map[string]bool
	Hp             int
}

func NewHangman(customWord string) *Hangman {
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

func (hangman *Hangman) IsAlive() bool {
	return hangman.Hp > 0
}

func (hangman *Hangman) IsWin() bool {
	for _, hasGuessed := range hangman.WordLetters {
		if !hasGuessed {
			return false
		}
	}
	return true
}

func (hangman *Hangman) CurrentWordStr() (wordStr string) {
	for _, letter := range strings.Split(hangman.ProtoWord, "") {
		if hasGuessed := hangman.WordLetters[letter]; hasGuessed {
			wordStr += letter
		} else {
			wordStr += "*"
		}
	}
	return
}
