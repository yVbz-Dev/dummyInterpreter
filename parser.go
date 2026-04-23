package main

import (
	"fmt"
)

func parser(tokens []Token) {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch {
		case token.Token == "print":
			var nextToken = nextToken(tokens, i)
			if nextToken.tokenType == "string" {
				fmt.Println(nextToken.Token)
			}
		}
	}
}

func nextToken(tokens []Token, currToken int) Token {
	if currToken >= len(tokens) {
		return Token{
			"EOF",
			0,
			0,
			"EOF",
		}
	}
	return tokens[currToken+1]
}
