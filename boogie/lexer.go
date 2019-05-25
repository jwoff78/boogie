package boogie

import "github.com/davecgh/go-spew/spew"

/*
CtxState defines a state machine we can use to parse lines of code.
*/
type CtxState int

const (
	// START sets an undefined state, used as a jumping off point.
	START CtxState = 0 + iota
	// LABEL sets the label state
	LABEL
)

/*
Lexeme is a type that structures a raw line of code string.
*/
type Lexeme struct {
	ID  string
	str string
}

/*
Lexer is a wrapper for all data and methods needed for converting lines of codes to Lexemes.
*/
type Lexer struct {
	ctxState  CtxState
	curCtx    Context
	curChar   string
	buf       string
	curLexeme Lexeme
	LCh       chan string
	delims    []string
	keys      []string
}

/*
NewLexer returns a reference to a new instance of Lexer.
*/
func NewLexer() *Lexer {
	return &Lexer{
		ctxState: START,
		LCh:      make(chan string),
		delims:   []string{":"},
		keys:     []string{"echo", "exit"},
	}
}

/*
Run a new instance of Lexer to receive lines of code over its channel.
*/
func (lexer *Lexer) Run() {
	for {
		select {
		case loc := <-lexer.LCh:
			// Convert the line of code to a Lexeme.
			for _, r := range loc {
				lexer.curChar = string(r)

				switch lexer.ctxState {
				case START:
					lexer.curCtx = Context{Start{}}
				case LABEL:
					lexer.curCtx = Context{Label{}}
				}

				// Append the new character to the transient buffer.
				lexer.buf += lexer.curCtx.getNextState(lexer)

				if lexer.curLexeme.ID != "" {
					spew.Dump(lexer.curLexeme)
				}
			}
		}
	}
}

func (lexer *Lexer) isDelim(str string) bool {
	return lexer.stringInSlice(str, lexer.delims)
}

func (lexer *Lexer) stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}

	return false
}
