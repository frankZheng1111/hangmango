package game

import (
	"fmt"
	"hangmango/models"
)

type Score struct {
	WinCount  int
	PlayCount int
}

func (score *Score) WinRate() (winRate float64) {
	if score.PlayCount == 0 {
		return 0
	}
	winRate = float64(score.WinCount) * 100
	winRate /= float64(score.PlayCount)
	return
}

func Play() {
	score := new(Score)
	fmt.Println("Welcome to hangmango!")
	printScore(*score)
	var finish string
	for {
		fmt.Printf("Input anything to guess a new word or end to finish the game: ")
		fmt.Scanln(&finish)
		if finish == "end" {
			break
		}
		startNewHangman(score)
	}
	printScore(*score)
	return
}

func startNewHangman(score *Score) {
	score.PlayCount++
	hangman := models.NewHangman("")
	fmt.Println(hangman.CurrentWordStr(), hangman.ProtoWord)
	var letter string
	for !hangman.IsWin() && hangman.IsAlive() {
		letter = ""
		fmt.Printf("You have %d more change. Please Guess a letter: ", hangman.Hp)
		fmt.Scanln(&letter)
		if err := hangman.Guess(letter); err != nil {
			fmt.Printf("\nWarning: %s!\n", err)
		}
		fmt.Printf("\nTARGET WORD: %s\n\n", hangman.CurrentWordStr())
	}
	if hangman.IsWin() {
		score.WinCount++
		fmt.Println("YOU WIN!")
	}
	return
}

func printScore(score Score) {
	fmt.Println("")
	fmt.Println("#############################################################################")
	fmt.Println("")
	fmt.Printf("YOU SCORE: | WIN: %d TIMES | PLAY: %d TIMES | RATE: %.2f %% \n", score.WinCount, score.PlayCount, score.WinRate())
	fmt.Println("")
	fmt.Println("#############################################################################")
}
