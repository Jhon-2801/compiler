package parser

import (
	"fmt"
	"os"

	"github.com/Jhon-2801/compiler/core/emitter"
	"github.com/Jhon-2801/compiler/core/lexer"
)

// Parser object keeps track of current token and checks if the code matches the grammar.
type Parser struct {
	symbols        map[string]struct{}
	labelsDeclared map[string]struct{}
	labelsGotoed   map[string]struct{}

	lexer     *lexer.Lexer
	emitter   *emitter.Emitter
	curToken  lexer.Token
	peekToken lexer.Token
}

// NewParser creates a new Parser instance with the provided lexer.
func NewParser(lexer *lexer.Lexer, emitter *emitter.Emitter) *Parser {
	parser := &Parser{
		lexer:          lexer,
		emitter:        emitter,
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
	p.emitter.HeaderLine("#include <stdio.h>")
	p.emitter.HeaderLine("int main(void){")
	// Since some newlines are required in our grammar, need to skip the excess.
	for p.checkToken(lexer.NEWLINE) {
		p.nextToken()
	}
	//Parse all the statements in the program.
	for !p.checkToken(lexer.EOF) {
		p.statement()
	}

	p.emitter.EmitLine("return 0;")
	p.emitter.EmitLine("}")

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
		p.nextToken()
		if p.checkToken(lexer.STRING) {
			//Simple string
			p.emitter.EmitLine("printf(\"" + p.curToken.Text + "\\n\");")
			p.nextToken()
		} else {
			p.emitter.Emit("printf(\"%" + ".2f\\n\", (float)(")
			p.expression()
			p.emitter.Emit("));")
		}

	} else if p.checkToken(lexer.IF) { //"IF" comparison "THEN" {statement} "ENDIF"
		p.nextToken()
		p.emitter.Emit("if(")
		p.comparison()

		p.match(lexer.TokenInfo{TokenType: lexer.THEN, Name: "THEN"})
		p.nl()
		p.emitter.EmitLine("){")
		//zero or more statements in the body.
		for !p.checkToken(lexer.ENDIF) {
			p.statement()
		}
		p.match(lexer.TokenInfo{TokenType: lexer.ENDIF, Name: "ENDIF"})
		p.emitter.EmitLine("}")

	} else if p.checkToken(lexer.WHILE) { // "WHILE" comparison "REPEAT" {statement} "ENDWHILE"
		p.nextToken()
		p.emitter.Emit("while(")
		p.comparison()

		p.match(lexer.TokenInfo{TokenType: lexer.REPEAT, Name: "REPEAT"})
		p.nl()
		p.emitter.EmitLine("){")

		//zero or more statements in the body.
		for !p.checkToken(lexer.ENDWHILE) {
			p.statement()
		}
		p.match(lexer.TokenInfo{TokenType: lexer.ENDWHILE, Name: "ENDWHILE"})
		p.emitter.EmitLine("}")

	} else if p.checkToken(lexer.LABEL) { // "LABEL" ident
		p.nextToken()
		if _, exists := p.labelsDeclared[p.curToken.Text]; exists {
			p.abort("Label already exists: " + p.curToken.Text)
		}
		p.labelsDeclared[p.curToken.Text] = struct{}{}

		p.emitter.EmitLine(p.curToken.Text + ":")
		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})

	} else if p.checkToken(lexer.GOTO) { // "GOTO" ident
		p.nextToken()
		p.labelsDeclared[p.curToken.Text] = struct{}{}
		p.emitter.EmitLine("goto" + p.curToken.Text + ":")
		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})

	} else if p.checkToken(lexer.LET) { // "LET" ident "=" expression
		p.nextToken()
		if _, exists := p.symbols[p.curToken.Text]; !exists {
			p.symbols[p.curToken.Text] = struct{}{}
			p.emitter.HeaderLine("float " + p.curToken.Text + ";")
		}
		p.emitter.Emit(p.curToken.Text + " = ")
		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})
		p.match(lexer.TokenInfo{TokenType: lexer.EQ, Name: "EQ"})
		p.expression()
		p.emitter.EmitLine(";")

	} else if p.checkToken(lexer.INPUT) { // "INPUT" ident
		p.nextToken()
		if _, exists := p.symbols[p.curToken.Text]; !exists {
			p.symbols[p.curToken.Text] = struct{}{}
			p.emitter.HeaderLine("float " + p.curToken.Text + ";")
		}

		p.emitter.EmitLine("if(0 == scanf(\"%" + "f\", &" + p.curToken.Text + ")) {")
		p.emitter.EmitLine(p.curToken.Text + " = 0;")
		p.emitter.Emit("scanf(\"%")
		p.emitter.EmitLine("*s\");")
		p.emitter.EmitLine("}")

		p.match(lexer.TokenInfo{TokenType: lexer.IDENT, Name: "IDENT"})
	} else { //    # This is not a valid statement. Error!
		p.abort("Invalid statement at " + p.curToken.Text + " (" + p.curToken.Kind.Name + ")")
	}

	p.nl()
}

// comparison ::= expression (("==" | "!=" | ">" | ">=" | "<" | "<=") expression)
func (p *Parser) comparison() {
	p.expression()

	if p.isComparisonOperator() {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
		p.expression()
	}
	for p.isComparisonOperator() {
		p.emitter.Emit(p.curToken.Text)
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
	p.term()
	//can have 0 or more +/- and expressions.
	for p.checkToken(lexer.PLUS) || p.checkToken(lexer.MINUS) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
		p.term()
	}
}

// term ::= unary {("/" | "*") unary}

func (p *Parser) term() {
	p.unary()
	//can have 0 or more *// and expressions.
	for p.checkToken(lexer.ASTERISK) || p.checkToken(lexer.SLASH) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
		p.unary()
	}
}

// unary ::= ["+" | "-"] primary

func (p *Parser) unary() {
	fmt.Println("UNARY")
	//can have 0 or more +/- and expressions.
	for p.checkToken(lexer.PLUS) || p.checkToken(lexer.MINUS) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
	}
	p.primary()
}

// primary ::= number | ident
func (p *Parser) primary() {
	if p.checkToken(lexer.NUMBER) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
	} else if p.checkToken(lexer.IDENT) {
		if _, exists := p.symbols[p.curToken.Text]; !exists {
			p.abort("Referencing variable before assignment: " + p.curToken.Text)
		}
		p.emitter.Emit(p.curToken.Text)
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
