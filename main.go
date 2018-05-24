package main

import (
	"fmt"
	"hangmango/models"
)

func main() {
	hangman := models.NewHangman("")
	fmt.Println(hangman.CurrentWordStr(), hangman.ProtoWord)
}
