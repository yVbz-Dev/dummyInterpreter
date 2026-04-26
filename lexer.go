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
		case char == '/' && line[i+1] == '/' && !readingToken:
			return tokens
		case isKeyword(string(char)) && currTokenType == "number":
			tokens = append(tokens, Token{currToken, lineNum, i, currTokenType})
			tokens = append(tokens, Token{string(char), lineNum, i + 1, "operator"})
			currTokenType = ""
			currToken = ""
			readingToken = false
		case readingToken:
			currToken += string(char)
		default:
			readingToken = true
			currTokenType = "command"
			currToken += string(char)
		}

		if i+1 >= len(line) {
			if currToken == "" || currToken == " " {
				continue
			}

			// append the fucking token
			tokens = append(tokens, Token{currToken, lineNum, i, currTokenType})
			currTokenType = ""
			currToken = ""
			readingToken = false
		}
	}

	// return tokens
	return tokens
}

func hasToken(tokens []Token, token Token) bool {
	var hasToken bool = false
	for _, t := range tokens {
		if t.Token == token.Token {
			hasToken = true
		}
	}
	return hasToken
}
