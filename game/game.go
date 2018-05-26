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
		fmt.Printf("TARGET WORD: %s You have %d more chance. Please Guess a letter: ", hangman.CurrentWordStr(), hangman.Hp)
		fmt.Scanln(&letter)
		if err := hangman.Guess(letter); err != nil {
			fmt.Printf("Warning: %s!\n", err)
		}
	}
	fmt.Printf("TARGET WORD: %s ANSWER: %s\n", hangman.CurrentWordStr(), hangman.ProtoWord)
	if hangman.IsWin() {
		score.WinCount++
		fmt.Println("YOU WIN!")
	} else {
		fmt.Println("YOU LOSE!")
	}
	return
}

func printScore(score Score) {
	fmt.Printf("\n#############################################################################\n\n")
	fmt.Printf("YOU SCORE: | WIN: %d TIMES | PLAY: %d TIMES | RATE: %.2f %% \n", score.WinCount, score.PlayCount, score.WinRate())
	fmt.Printf("\n#############################################################################\n\n")
}
