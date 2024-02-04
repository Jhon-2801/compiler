package parser

import (
	"fmt"
	"os"

	"github.com/Jhon-2801/compiler/core/lexer"
)

// Parser object keeps track of current token and checks if the code matches the grammar.
type Parser struct {
	symbols        map[string]struct{}
	labelsDeclared map[string]struct{}
	labelsGotoed   map[string]struct{}

	lexer     *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

// NewParser creates a new Parser instance with the provided lexer.
func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:          lexer,
		symbols:        make(map[string]struct{}),
		labelsDeclared: make(map[string]struct{}),
		labelsGotoed:   make(map[string]struct{}),
	}
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
	// Since some newlines are required in our grammar, need to skip the excess.
	for p.checkToken(lexer.NEWLINE) {
		p.nextToken()
	}
	//Parse all the statements in the program.
	for !p.checkToken(lexer.EOF) {
		p.statement()
	}
	for label := range p.labelsGotoed {
		if _, exists := p.labelsDeclared[label]; !exists {
			p.abort("Attempting to GOTO to undeclared label: " + label)
		}
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

	} else if p.checkToken(lexer.IF) { //"IF" comparison "THEN" {statement} "ENDIF"
		fmt.Println("STATEMENT-IF")
		p.nextToken()
		p.comparison()

		p.match(lexer.TokenInfo{TokenType: lexer.THEN, Name: "THEN"})
		p.nl()

		//zero or more statements in the body.
		for !p.checkToken(lexer.ENDIF) {
			p.statement()
		}
		p.match(lexer.TokenInfo{TokenType: lexer.ENDIF, Name: "ENDIF"})

	} else if p.checkToken(lexer.WHILE) { // "WHILE" comparison "REPEAT" {statement} "ENDWHILE"
		fmt.Println("STATEMENT-WHILE")
		p.nextToken()
		p.comparison()

		p.match(lexer.TokenInfo{TokenType: lexer.REPEAT, Name: "REPEAT"})
		p.nl()

		//zero or more statements in the body.
		for !p.checkToken(lexer.ENDWHILE) {
			p.statement()
		}
		p.match(lexer.TokenInfo{TokenType: lexer.ENDWHILE, Name: "ENDWHILE"})

	} else if p.checkToken(lexer.LABEL) { // "LABEL" ident
		fmt.Println("STATEMENT-LABEL")
		p.nextToken()
		if _, exists := p.labelsDeclared[p.curToken.Text]; exists {
			p.abort("Label already exists: " + p.curToken.Text)
		}
		p.labelsDeclared[p.curToken.Text] = struct{}{}
		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})
	} else if p.checkToken(lexer.GOTO) { // "GOTO" ident
		fmt.Println("STATEMENT-GOTO")
		p.nextToken()

		p.labelsDeclared[p.curToken.Text] = struct{}{}
		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})
	} else if p.checkToken(lexer.LET) { // "LET" ident "=" expression
		fmt.Println("STATEMENT-LET")
		p.nextToken()
		if _, exists := p.symbols[p.curToken.Text]; !exists {
			p.symbols[p.curToken.Text] = struct{}{}
		}
		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})
		p.match(lexer.TokenInfo{TokenType: lexer.EQ, Name: "EQ"})
		p.expression()

	} else if p.checkToken(lexer.INPUT) { // "INPUT" ident
		fmt.Println("STATEMENT-INPUT")
		p.nextToken()
		if _, exists := p.symbols[p.curToken.Text]; !exists {
			p.symbols[p.curToken.Text] = struct{}{}
		}
		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})
	} else { //    # This is not a valid statement. Error!
		p.abort("Invalid statement at " + p.curToken.Text + " (" + p.curToken.Kind.Name + ")")
	}

	p.nl()
}

// comparison ::= expression (("==" | "!=" | ">" | ">=" | "<" | "<=") expression)
func (p *Parser) comparison() {
	fmt.Println("COMPARISON")

	p.expression()

	if p.isComparisonOperator() {
		p.nextToken()
		p.expression()
	} else {
		p.abort("Expected comparison operator at: " + p.curToken.Text)
	}

	for p.isComparisonOperator() {
		p.nextToken()
		p.expression()
	}
}

// return true if the current token is a comparison operator
func (p *Parser) isComparisonOperator() bool {
	return p.checkToken(lexer.GT) || p.checkToken(lexer.GTEQ) || p.checkToken(lexer.LT) || p.checkToken(lexer.LTEQ) || p.checkToken(lexer.EQEQ) || p.checkToken(lexer.NOTEQ)
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

	p.unary()

	//can have 0 or more *// and expressions.
	for p.checkToken(lexer.ASTERISK) || p.checkToken(lexer.SLASH) {
		p.nextToken()
		p.unary()
	}
}

// unary ::= ["+" | "-"] primary

func (p *Parser) unary() {
	fmt.Println("UNARY")

	//can have 0 or more +/- and expressions.
	for p.checkToken(lexer.PLUS) || p.checkToken(lexer.MINUS) {
		p.nextToken()
	}
	p.primary()
}

// primary ::= number | ident
func (p *Parser) primary() {
	fmt.Println("PRIMARY (" + p.curToken.Text + ")")

	if p.checkToken(lexer.NUMBER) {
		p.nextToken()
	} else if p.checkToken(lexer.IDENT) {
		if _, exists := p.symbols[p.curToken.Text]; !exists {
			p.abort("Referencing variable before assignment: " + p.curToken.Text)
		}
		p.nextToken()
	} else {
		//Error!
		p.abort("Unexpected token at " + p.curToken.Text)
	}
}

// NewLine
func (p *Parser) nl() {
	fmt.Println("NEWLINE")

	p.match(lexer.TokenInfo{TokenType: lexer.NEWLINE, Name: "NEWLINE"})

	for p.checkToken(lexer.NEWLINE) {
		p.nextToken()
	}
}
