package models

import "testing"

func TestHangmanIsAlive(t *testing.T) {
	hangman := new(Hangman)
	hangman.Hp = 0
	isAlive := hangman.IsAlive()
	if isAlive {
		t.Error("TestHangmanIsAlive fail")
	}
}
