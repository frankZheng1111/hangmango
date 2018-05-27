package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinRateReturnZero(t *testing.T) {
	score := new(Score)
	assert.Equal(t, score.WinRate(), float64(0))
}
