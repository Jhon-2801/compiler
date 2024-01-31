package main

import (
	"fmt"

	"github.com/Jhon-2801/compiler/core/lexer"
)

func main() {

	source := "+-123 9.8654*/"
	lex := lexer.NewLexer(source)

	token := lex.GetToken()
	for token.Kind != lexer.EOF {
		fmt.Println(token.Text)
		token = lex.GetToken()
	}
}
