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
	parser := &Parser{lexer: lexer}
	parser.nextToken()
	parser.nextToken() // Call this twice to initialize current and peek.
	return parser
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
	p.nextToken()
}

// Advances the current token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.GetToken()
	//No need to worry about passing the EOF, lexer handles that.
}

//Production rules.

// program ::= {statement}
func (p *Parser) Program() {
	fmt.Println("PROGRAM")
	//Parse all the statements in the program.
	for !p.checkToken(lexer.EOF) {

	}
}
func (p *Parser) abort(message string) {
	fmt.Printf("Error. %s\n", message)
	os.Exit(1)
}

//One of the following statements

func (p *Parser) statement() {
	//Check the first token to see what kind of statement this is.

	//"PRINT" (expression | string)
	if p.checkToken(lexer.PRINT) {
		print("STATEMENT-PRINT")
		p.nextToken()

		if p.checkToken(lexer.STRING) {
			//Simple string
			p.nextToken()
		} else {

		}
	}
}
