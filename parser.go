package main

import (
	"fmt"
	"sort"
	"strconv"
)

var equationForce map[string]int = map[string]int{
	KW_PLUS:  1,
	KW_MINUS: 1,
	KW_MULT:  2,
	KW_DIV:   2,
}

func parser(tokens []Token) {
	// var program []ASTNode
	var pos int = 0
	var program []ASTNode
	for pos < len(tokens) {
		token := tokens[pos]

		switch {
		case token.Token == KW_PRINT:
			// vars
			var printVal string

			// get all the + signs and the tokens
			var posUpdaterVar int = 0
			for i := pos + 1; i < len(tokens); i++ {
				var iToken Token = tokens[i]
				if isKeyword(iToken.Token) && iToken.Token != KW_PLUS {
					break
				} else {
					if iToken.Token == KW_PLUS {
						continue
					}
					VarInMemory := memory["var_"+iToken.Token]
					if VarInMemory != nil {
						printVal += VarInMemory.(string)
					} else {
						printVal += iToken.Token
					}
				}
				i++
				posUpdaterVar++
			}

			// create node
			var node ASTNode
			node.NodeAction = token.Token
			node.NodeArgs = map[string]any{
				"value": printVal,
			}
			pos += posUpdaterVar
			program = append(program, node)
		case token.Token == KW_VAR:
			// vars
			var varName Token = nextToken(tokens, pos)
			var equalSign Token = nextToken(tokens, pos+1)
			var value Token = nextToken(tokens, pos+2)
			var valueStr string

			// get all the + signs and the tokens
			var posUpdateVar = 0
			for i := pos; i < len(tokens); i++ {
				// var
				var iToken Token = tokens[i]

				if isKeyword(iToken.Token) && iToken.Token != KW_PLUS {
					break
				} else {
					if iToken.Token == KW_PLUS {
						continue
					}
					varInMemory := memory["var_"+iToken.Token]
					if varInMemory != nil {
						valueStr += varInMemory.(string)
					} else {
						valueStr += iToken.Token
					}
				}
				// update pos
				posUpdateVar++
				i++
			}

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
				"value":   valueStr,
				"varName": varName.Token,
			}
			program = append(program, node)
			pos += posUpdateVar
		case token.tokenType == "number":
			var equation = []Token{}
			var posUpdaterVar = 0
			for i := pos; i < len(tokens); i++ {
				var iToken Token = tokens[i]
				if isKeyword(iToken.Token) && !isOperator(iToken.Token) {
					break
				}
				equation = append(equation, iToken)
				posUpdaterVar++
			}

			// calculate the equation
			// var result int = 0
			var onHold []int = []int{}
			for i := 0; i < len(equation); i++ {
				var iToken Token = equation[i]
				if !isOperator(iToken.Token) {
					continue
				}
				onHold = append(onHold, i)
			}

			sort.Slice(onHold, func(i, j int) bool {
				return i > j
			})
			fmt.Println(equation)
			fmt.Println(onHold)

			pos += posUpdaterVar
		}
		pos++
		runner(program)
	}
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

func isOperator(s string) bool {
	isOperator := s == KW_PLUS || s == KW_MINUS || s == KW_MULT || s == KW_DIV
	return isOperator
}

func operatorForce(s string) int {
	return equationForce[s]
}
