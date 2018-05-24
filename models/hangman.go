package models

import (
	"fmt"
	"hangmango/config"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var dictionary []string

func init() {
	path := config.GameSettingConfig.DictionaryPath
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic("missing dictionary_path in config")
	}
	for _, letter := range strings.Split(string(content), "\n") {
		dictionary = append(dictionary, letter)
	}
}

type Hangman struct {
	ProtoWord      string
	GuessedLetters map[string]bool
	WordLetters    map[string]bool
	Hp             int
}

func NewHangman(customWord string) *Hangman {
	hangman := new(Hangman)
	if customWord == "" {
		source := rand.NewSource(time.Now().Unix())
		randMachine := rand.New(source)
		randIndex := randMachine.Intn(len(dictionary) - 1)
		fmt.Println(randIndex)
		customWord = dictionary[randIndex]
	}
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
