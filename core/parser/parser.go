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
		p.statement()
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
		fmt.Println("STATEMENT-PRINT")
		p.nextToken()

		if p.checkToken(lexer.STRING) {
			//Simple string
			p.nextToken()
		} else {
			p.expression()
		}
	}
	//Newline
	p.nl()
}

//expression ::= term {("-" | "+") term}

func (p *Parser) expression() {
	fmt.Println("EXPRESSION")

	p.term()
	//can have 0 or more +/- and expressions.
	for p.checkToken(lexer.PLUS) || p.checkToken(lexer.MINUS) {
		p.nextToken()
		p.term()
	}
}

// term ::= unary {("/" | "*") unary}

func (p *Parser) term() {
	fmt.Println("TERM")

	//can have 0 or more +/- and expressions.
	for p.checkToken(lexer.ASTERISK) || p.checkToken(lexer.SLASH) {
		p.nextToken()
	}
}

// unary ::= ["+" | "-"] primary

func (p *Parser) unary() {
	fmt.Println("UNARY")

	//can have 0 or more +/- and expressions.
	for p.checkToken(lexer.PLUS) || p.checkToken(lexer.MINUS) {
		p.nextToken()
	}

}

func (p *Parser) nl() {
	fmt.Println("NEWLINE")

	p.match(lexer.TokenInfo{TokenType: lexer.NEWLINE, Name: "NEWLINE"})

	for p.checkToken(lexer.NEWLINE) {
		p.nextToken()
	}
}
