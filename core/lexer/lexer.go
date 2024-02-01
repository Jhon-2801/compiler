package lexer

import (
	"log"
)

type Lexer struct {
	Source  string
	CurChar string
	CurPos  int
}

type Token struct {
	Text string    // The token's actual text. Used for identifiers, strings, and numbers.
	Kind TokenType // The TokenType that this token is classified as.
}

// TokenType represents the type of a token.
type TokenType int

const (
	EOF     TokenType = -1
	NEWLINE TokenType = 0
	NUMBER  TokenType = 1
	IDENT   TokenType = 2
	STRING  TokenType = 3
	// Keywords.
	LABEL    TokenType = 101
	GOTO     TokenType = 102
	PRINT    TokenType = 103
	INPUT    TokenType = 104
	LET      TokenType = 105
	IF       TokenType = 106
	THEN     TokenType = 107
	ENDIF    TokenType = 108
	WHILE    TokenType = 109
	REPEAT   TokenType = 110
	ENDWHILE TokenType = 111
	// Operators.
	EQ       TokenType = 201
	PLUS     TokenType = 202
	MINUS    TokenType = 203
	ASTERISK TokenType = 204
	SLASH    TokenType = 205
	EQEQ     TokenType = 206
	NOTEQ    TokenType = 207
	LT       TokenType = 208
	LTEQ     TokenType = 209
	GT       TokenType = 210
	GTEQ     TokenType = 211
)

//Constructor

func NewLexer(source string) *Lexer {
	lexer := &Lexer{
		Source:  source + "\n",
		CurChar: "",
		CurPos:  -1,
	}
	lexer.NextChar()
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
	log.Fatal("lexing error." + message)
}

// Skip whitespace except newlines.
func (l *Lexer) SkipWhitespace() {
	for l.CurChar == " " || l.CurChar == "\t" || l.CurChar == "\r" {
		l.NextChar()
	}
}

// Skip comments in the code.
func (l *Lexer) SkipComments() {
	if l.CurChar == "#" {
		for l.CurChar != "\n" {
			l.NextChar()
		}
	}
}

// Check whether this token is a number
func (l *Lexer) Isdigit() bool {
	if l.Peek() >= "0" && l.Peek() <= "9" {
		return true
	}
	return false
}

// Return the next token.
func (l *Lexer) GetToken() Token {
	l.SkipWhitespace()
	l.SkipComments()
	var token Token
	var lastChar string
	switch l.CurChar {
	case "+":
		token = Token{l.CurChar, PLUS}
	case "-":
		token = Token{l.CurChar, MINUS}
	case "*":
		token = Token{l.CurChar, ASTERISK}
	case "/":
		token = Token{l.CurChar, SLASH}
	case "=":
		//Check whether this token is = or ==
		if l.Peek() == "=" {
			lastChar = l.CurChar
			l.NextChar()
			token = Token{lastChar + l.CurChar, EQEQ}
		} else {
			token = Token{l.CurChar, EQ}
		}
	case ">":
		//Check whether this token is > or >=
		if l.Peek() == "=" {
			lastChar = l.CurChar
			l.NextChar()
			token = Token{lastChar + l.CurChar, GTEQ}
		} else {
			token = Token{l.CurChar, GT}
		}
	case "<":
		//Check whether this token is < or <=
		if l.Peek() == "=" {
			lastChar = l.CurChar
			l.NextChar()
			token = Token{lastChar + l.CurChar, LTEQ}
		} else {
			token = Token{l.CurChar, LT}
		}
	case "!":
		//Check whether this token is ! or !=
		if l.Peek() == "=" {
			lastChar = l.CurChar
			l.NextChar()
			token = Token{lastChar + l.CurChar, NOTEQ}
		} else {
			l.Abort("Exected !=, got !" + l.Peek())
		}
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		//Leading character is a digit, so this must be a number.
		//Get all consecutive digits and decimal if there is one.
		startPos := l.CurPos

		for l.Isdigit() {
			l.NextChar()
		}
		// Decimal
		if l.Peek() == "." {
			l.NextChar()

			if !l.Isdigit() {
				l.Abort("Ilegal character in number")
			}
			for l.Isdigit() {
				l.NextChar()
			}
		}
		tokText := l.Source[startPos : l.CurPos+1] // Get the substring.
		token = Token{tokText, NUMBER}
	case "\"":
		//Get characters between quotation.
		l.NextChar()
		startPos := l.CurPos

		for l.CurChar != "\"" {
			if l.CurChar == "\r" || l.CurChar == "\n" || l.CurChar == "\t" || l.CurChar == "\\" || l.CurChar == "%" {
				l.Abort("Ilegal character in string")
			}
			l.NextChar()
		}
		tokText := l.Source[startPos:l.CurPos] // Get the substring
		token = Token{tokText, STRING}
	case "\n":
		token = Token{l.CurChar, NEWLINE}
	case "":
		token = Token{Text: "", Kind: EOF}
	default:
		l.Abort("Unknow token: " + l.CurChar)
		// Unknown token!
	}
	l.NextChar()
	return token
}
