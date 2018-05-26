package game

import "testing"

func TestWinRateReturnZero(t *testing.T) {
	score := new(Score)
	if score.WinRate() != 0 {
		t.Error("TestWinRateReturnZero fail")
	}
}
