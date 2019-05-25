package boogie

/*
Context is a wrapper for all data and method needed to get the current context.
*/
type Context struct {
	handler interface {
		getNextState(lexer *Lexer) string
	}
}

func (ctx *Context) getNextState(lexer *Lexer) string {
	return ctx.handler.getNextState(lexer)
}
