package boogie

/*
AST wraps the Bastract Syntax Tree.
*/
type AST struct {
	leaves []AST
}

/*
Parser is a wrapper for all data and methods related to parsing Lexemes into an AST.
*/
type Parser struct {
	PCh chan *Lexeme
	ast AST
}

/*
NewParser returns a reference to a new instance of Parser.
*/
func NewParser() *Parser {
	return &Parser{
		PCh: make(chan *Lexeme),
	}
}

/*
Run a new instance of Parser to receive Lexemes over its channel.
*/
func (parser *Parser) Run() {
	for {
		select {
		case lexeme := <-parser.PCh:
			// Dome some...
			err := parser.parse(lexeme)

			if err != nil {
				panic(err)
			}
		}
	}
}

func (parser *Parser) parse(lexeme *Lexeme) error {
	return nil
}
