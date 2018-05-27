package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHangmanIsAlive(t *testing.T) {
	hangman := new(Hangman)
	hangman.Hp = 0
	hangman.IsAlive()
	assert.False(t, hangman.IsAlive())
	hangman.Hp = 1
	assert.True(t, hangman.IsAlive())
}
