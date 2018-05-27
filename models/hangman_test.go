package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHangmanWithRandomWord(t *testing.T) {
	hangman := NewHangman("")
	assert.Equal(t, hangman.ProtoWord, "abandon")
	assert.Equal(t, hangman.Hp, 10)
	assert.Equal(t, hangman.GuessedLetters, make(map[string]bool))
	wordLetters := map[string]bool{
		"a": false,
		"b": false,
		"n": false,
		"o": false,
		"d": false,
	}
	assert.Equal(t, hangman.WordLetters, wordLetters)
}

func TestNewHangmanWithCustomWord(t *testing.T) {
	hangman := NewHangman("test")
	assert.Equal(t, hangman.ProtoWord, "test")
	assert.Equal(t, hangman.Hp, 10)
	assert.Equal(t, hangman.GuessedLetters, make(map[string]bool))
	wordLetters := map[string]bool{
		"e": false,
		"t": false,
		"s": false,
	}
	assert.Equal(t, hangman.WordLetters, wordLetters)
}

func TestHangmanIsAlive(t *testing.T) {
	hangman := new(Hangman)
	hangman.Hp = 0
	hangman.IsAlive()
	assert.False(t, hangman.IsAlive())
	hangman.Hp = 1
	assert.True(t, hangman.IsAlive())
}

func TestHangmanIsWin(t *testing.T) {
	hangman := new(Hangman)
	hangman.Hp = 0
	assert.False(t, hangman.IsWin())
	hangman.Hp = 1
	hangman.WordLetters = map[string]bool{"a": true}
	assert.True(t, hangman.IsWin())
	hangman.WordLetters["b"] = false
	assert.False(t, hangman.IsWin())
}

func TestHangmanCurrentWordStr(t *testing.T) {
	hangman := new(Hangman)
	hangman.ProtoWord = "abab"
	hangman.WordLetters = map[string]bool{"a": true, "b": false}
	assert.Equal(t, hangman.CurrentWordStr(), "a*a*")
}

func TestHangmanGuess(t *testing.T) {
	hangman := NewHangman("test")
	hangman.Hp = 0
	assert.NotNil(t, hangman.Guess("t"))
	hangman.Hp = 10
	assert.NotNil(t, hangman.Guess("ta"))
	assert.Equal(t, hangman.Hp, 9)
	assert.NotNil(t, hangman.Guess(""))
	assert.Equal(t, hangman.Hp, 8)
	hangman.Guess("t")
	assert.True(t, hangman.GuessedLetters["t"])
	assert.True(t, hangman.WordLetters["t"])
	assert.Equal(t, hangman.Hp, 8)
	hangman.Guess("a")
	assert.Equal(t, hangman.Hp, 7)
	assert.True(t, hangman.GuessedLetters["a"])
	assert.NotNil(t, hangman.Guess("a"))
	assert.Equal(t, hangman.Hp, 6)
}
