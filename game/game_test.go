package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinRateReturnZero(t *testing.T) {
	score := new(Score)
	assert.Equal(t, score.WinRate(), float64(0))
}

func TestWinRateReturn(t *testing.T) {
	score := new(Score)
	score.PlayCount = 10
	score.WinCount = 6
	assert.Equal(t, score.WinRate(), float64(60))
}

func TestPrintScore(t *testing.T) {
	score := new(Score)
	printScore(*score)
}

func TestStartNewHangman(t *testing.T) {
	score := new(Score)
	startNewHangman(score)
	assert.Equal(t, score.PlayCount, 1)
}
