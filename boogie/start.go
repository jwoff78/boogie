package boogie

/*
Start handles the start context.
*/
type Start struct{}

func (start Start) getNextState(lexer *Lexer) string {
	if lexer.ctxState == START && lexer.isDelim(lexer.curChar) {
		lexer.ctxState = LABEL
		return ""
	} else if lexer.isKey(lexer.buf) {
		// We have a keyword.
		lexer.curLexeme.ID = lexer.buf
		lexer.ctxState = KEYWORD
	} else {
		// lexer.buf = ""
		lexer.curLexeme = Lexeme{}
	}

	return lexer.curChar
}
