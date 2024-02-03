package parser

import (
	"fmt"
	"os"

	"github.com/Jhon-2801/compiler/core/lexer"
)

// Parser object keeps track of current token and checks if the code matches the grammar.
type Parser struct {
	lexer *lexer.Lexer
}

// NewParser creates a new Parser instance with the provided lexer.
func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

// Return true if the current token matches.
func (p *Parser) CheckToken(kind lexer.TokenType) bool {
	return false
}

// Try to match current token. If not, error. Advances the current token.
func (p *Parser) Match(Kind lexer.TokenType) {

}

// Advances the current token
func (p *Parser) nextToken(Kind lexer.TokenType) {

}

func (p *Parser) abort(message string) {
	fmt.Printf("Error. %s\n", message)
	os.Exit(1)
}
