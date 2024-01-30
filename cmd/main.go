package main

import (
	"fmt"

	"github.com/Jhon-2801/compiler/lexer"
)

func main() {

	source := "LET foobar = 123"
	lex := lexer.NewLexer(source)

	for lex.Peek() != "" {
		fmt.Println(lex.CurChar)
		lex.NextChar()
	}
}
