package main

import (
	"fmt"

	"github.com/Jhon-2801/compiler/core/lexer"
)

func main() {

	source := "IF+-123 foo*THEN/"
	lex := lexer.NewLexer(source)

	token := lex.GetToken()
	for token.Kind != lexer.EOF {
		fmt.Println(token.Kind)
		token = lex.GetToken()
	}
}
