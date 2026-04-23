package main

import (
	"bufio"
	"fmt"
)

type Token struct {
	Token     string
	Line      int
	Column    int
	tokenType string
}

func lexer(sourceCode bufio.Scanner) {
	var lineNum int = 0
	for sourceCode.Scan() {
		// get line
		lineNum++
		line := sourceCode.Text()

		// tokenize!
		var tokens []Token = tokenize(line, lineNum)
		for _, token := range tokens {
			fmt.Println(token.Token)
		}
	}
}

func tokenize(line string, lineNum int) []Token {
	// vars
	var currToken string
	var readingToken bool = false
	var tokens = []Token{}

	// read trought the line!
	for i := 0; i < len(line); i++ {
		char := line[i]

		switch {
		case char == ' ':
			continue
		case char == '"':
			if readingToken {
				readingToken = false
				tokens = append(tokens, Token{currToken, lineNum, i, "string"})
				currToken = ""
			} else {
				readingToken = true
			}
		case readingToken:
			currToken += string(char)
		default:
			if readingToken {
				readingToken = false
				tokens = append(tokens, Token{currToken, lineNum, i, "command"})
			} else {
				readingToken = true
				currToken += string(char)
			}

		}
	}

	// return tokens
	return tokens
}
