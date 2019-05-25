package boogie

/*
Label handles the label context.
*/
type Label struct{}

func (label Label) getNextState(lexer *Lexer) string {
	if lexer.ctxState == LABEL {
		lexer.curLexeme.ID = "label"
		lexer.curLexeme.str = lexer.buf
		lexer.ctxState = START
	}

	return lexer.curChar
}
