package main

import (
	"bufio"
)

func lexer(sourceCode bufio.Scanner) {
	var lineNum int = 0
	for sourceCode.Scan() {
		// get line
		lineNum++
		line := sourceCode.Text()

		// tokenize!
		var tokens []Token = tokenize(line, lineNum)
		parser(tokens)
	}
}

func tokenize(line string, lineNum int) []Token {
	// vars
	var currToken string
	var currTokenType string
	var readingToken bool = false
	var tokens = []Token{}

	// read trought the line!
	for i := 0; i < len(line); i++ {
		char := line[i]

		// cases
		switch {
		case char == ' ':
			if readingToken && currTokenType != "string" {
				tokens = append(tokens, Token{currToken, lineNum, i, currTokenType})
				currTokenType = ""
				currToken = ""
				readingToken = false
			} else if readingToken {
				currToken += string(char)
			}
		case char == '"':
			if readingToken {
				readingToken = false
				tokens = append(tokens, Token{currToken, lineNum, i, currTokenType})
				currToken = ""
				currTokenType = ""
			} else {
				readingToken = true
				currTokenType = "string"
			}
		case (char >= '0' && char <= '9') && currTokenType != "string":
			readingToken = true
			currTokenType = "number"
			currToken += string(char)
		case readingToken:
			currToken += string(char)
		default:
			readingToken = true
			currTokenType = "command"
			currToken += string(char)

		}
	}

	// return tokens
	return tokens
}
