package lexer

type Lexer struct {
	Source  string
	CurChar string
	CurPos  int
}

//Constructor

func NewLexer(source string) *Lexer {
	lexer := &Lexer{
		Source:  source + "\n",
		CurChar: "",
		CurPos:  -1,
	}
	return lexer
}

// Process the next character.
func (l *Lexer) NextChar() {
	if l.CurPos+1 < len(l.Source) {
		l.CurPos++
		l.CurChar = string(l.Source[l.CurPos])
	} else {
		l.CurChar = "" //End of the source code is reached, set curChar to an empty string to indicate End of File
	}
}

// Return the lookahead character.
func (l *Lexer) Peek() string {
	if l.CurPos+1 >= len(l.Source) {
		return ""
	}
	return string(l.Source[l.CurPos+1])
}

// Invalid token found, print error message and exit.
func (l *Lexer) Abort(message string) {

}

// Skip whitespace except newlines.
func (l *Lexer) SkipWhitespace() {

}

// Skip comments in the code.
func (l *Lexer) GetToken() string {
	return ""
}
