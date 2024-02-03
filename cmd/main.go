package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Teeny Tiny Compiler")

	if len(os.Args) != 2 {
		fmt.Println("Error: Compiler needs sourse file as argument.")
		os.Exit(1)
	}

	fileName := os.Args[1]
	_, err := readFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", fileName, err)
		os.Exit(1)
	}
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
