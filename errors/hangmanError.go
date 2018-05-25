package errors

var MsgTexts map[string]string

func init() {
	MsgTexts = make(map[string]string)
	MsgTexts["input.guessed.letter"] = "You has guessed this letter"
	MsgTexts["input.nothing"] = "You need to guess a letter"
	MsgTexts["input.multi.letter"] = "You can only guess one letter one time"
}

type HangmanError struct {
	msg string
}

func (hangmanError *HangmanError) Error() string {
	if msgText, ok := MsgTexts[hangmanError.msg]; ok {
		return msgText
	} else {
		return "Unknown Error"
	}
}

func NewHangmanError(msg string) *HangmanError {
	hangmanError := new(HangmanError)
	hangmanError.msg = msg
	return hangmanError
}
