package main

import (
	"fmt"
	"strconv"
)

func parser(tokens []Token) {
	// var program []ASTNode
	var pos int = 0
	var program []ASTNode
	for pos < len(tokens) {
		token := tokens[pos]

		switch {
		case token.Token == KW_PRINT:
			// create node
			var node ASTNode
			node.NodeAction = token.Token
			node.NodeArgs = map[string]any{
				"value": nextToken(tokens, pos).Token,
			}
			pos++
			program = append(program, node)
		case token.Token == KW_VAR:
			// vars
			var varName Token = nextToken(tokens, pos)
			var equalSign Token = nextToken(tokens, pos+1)
			var value Token = nextToken(tokens, pos+2)

			// check
			if isKeyword(varName.Token) {
				fmt.Println("Syntax error: variable name cannot be a keyword at line " + strconv.Itoa(varName.Line) + " column " + strconv.Itoa(varName.Column))
				return
			}
			if equalSign.Token != "=" {
				fmt.Println("Syntax error: expected '=' after variable name at line " + strconv.Itoa(equalSign.Line) + " column " + strconv.Itoa(equalSign.Column))
				return
			}
			if value.Token == ("EOF") {
				fmt.Println("Syntax error: expected a value after '=' at line " + strconv.Itoa(value.Line) + " column " + strconv.Itoa(value.Column))
				return
			}

			// create node
			var node ASTNode
			node.NodeAction = token.Token
			node.NodeArgs = map[string]any{
				"value":   value.Token,
				"varName": varName.Token,
			}
			program = append(program, node)
			pos += 2
		}
		pos++
	}
	runner(program)
}

func nextToken(tokens []Token, currToken int) Token {
	if currToken+1 > len(tokens)-1 {
		return Token{
			"EOF",
			0,
			0,
			"EOF",
		}
	}
	return tokens[currToken+1]
}

func prevToken(tokens []Token, currToken int) Token {
	if currToken-1 < 0 {
		return Token{
			"EOF",
			0,
			0,
			"EOF",
		}
	}
	return tokens[currToken-1]
}

func isKeyword(s string) bool {
	_, exists := keywords[s]
	return exists
}
