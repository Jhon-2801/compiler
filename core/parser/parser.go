package parser

import (
	"fmt"
	"os"

	"github.com/Jhon-2801/compiler/core/lexer"
)

// Parser object keeps track of current token and checks if the code matches the grammar.
type Parser struct {
	lexer     *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

// NewParser creates a new Parser instance with the provided lexer.
func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

// Return true if the current token matches.
func (p *Parser) checkToken(kind lexer.TokenType) bool {
	return kind == p.curToken.Kind.TokenType
}

// Return true if the next token matches.
func (p *Parser) checkPeek(kind lexer.TokenType) bool {
	return kind == p.curToken.Kind.TokenType
}

// Try to match current token. If not, error. Advances the current token.
func (p *Parser) match(Kind lexer.TokenInfo) {
	if !p.checkPeek(Kind.TokenType) {
		p.abort("Expected" + Kind.Name + ", got " + p.curToken.Kind.Name)
	}
}

// Advances the current token
func (p *Parser) nextToken(Kind lexer.TokenType) {

}

func (p *Parser) abort(message string) {
	fmt.Printf("Error. %s\n", message)
	os.Exit(1)
}
