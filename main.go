package main

import (
	"fmt"
	"hangmango/models"
)

func main() {
	hangman := models.NewHangman("test")
	fmt.Println(hangman)
}
