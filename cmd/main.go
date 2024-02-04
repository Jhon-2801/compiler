package main

import (
	"fmt"
	"os"

	"github.com/Jhon-2801/compiler/core/emitter"
	"github.com/Jhon-2801/compiler/core/lexer"
	"github.com/Jhon-2801/compiler/core/parser"
)

func main() {
	fmt.Println("Teeny Tiny Compiler")

	if len(os.Args) != 2 {
		fmt.Println("Error: Compiler needs sourse file as argument.")
		os.Exit(1)
	}

	fileName := os.Args[1]
	sourse, err := readFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", fileName, err)
		os.Exit(1)
	}
	lexer := lexer.NewLexer(sourse)
	emit := emitter.NewEmitter("out.c")
	parser := parser.NewParser(lexer, emit)

	parser.Program() //Start the parser
	emit.WriteFile()
	fmt.Println("Parsing completed")
}

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	size := stat.Size()

	content := make([]byte, size)

	_, err = file.Read(content)

	if err != nil {
		return "", err
	}

	return string(content), err
}
