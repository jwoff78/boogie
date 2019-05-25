package main

import (
	"bufio"
	"os"

	"github.com/theapemachine/boogie/boogie"
)

func main() {
	lexer := boogie.NewLexer()
	go lexer.Run()

	file, err := os.Open("hello.boo")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Looping over the file line by line.
		lexer.LCh <- scanner.Text()
	}
}
