package main

import (
	"fmt"
	"slices"
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
			var posUpdaterVar int = 2
			for i := pos + 1; i < len(tokens); i++ {
				var iToken Token = tokens[i]
				if isKeyword(iToken.Token) && iToken.Token != KW_PLUS {
					break
				} else {
					if iToken.Token == KW_PLUS {
						continue
					}

					// check equation
					VarInMemory := memory["var_"+iToken.Token]
					var isVar bool = false
					if VarInMemory != nil && iToken.tokenType == "number" {
						_, err := strconv.Atoi(VarInMemory.(string))
						if err != nil {
							return
						}
						isVar = true
					}

					if iToken.tokenType == "number" || isVar {
						var equation []Token = []Token{}
						var result int
						var iRemover int = 2
						if isVar {
							iRemover = 0
						}

						for j := i - iRemover; j < len(tokens); j++ {
							var jToken Token = tokens[j]
							if isKeyword(jToken.Token) && !isOperator(jToken.Token) {
								break
							}

							VarInMemory := memory["var_"+jToken.Token]
							if VarInMemory != nil {
								equation = append(equation, Token{VarInMemory.(string), jToken.Line, "number"})
							} else {
								equation = append(equation, jToken)
							}

						}

						result = solveEquation(equation)
						printVal = strconv.Itoa(result)
						pos += len(equation)
						break
					} else {
						VarInMemory := memory["var_"+iToken.Token]
						if VarInMemory != nil && iToken.tokenType == "number" {
							printVal += VarInMemory.(string)
						} else {
							printVal += iToken.Token
						}
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
			var posUpdateVar = 1
			for i := pos + 3; i < len(tokens); i++ {
				// var
				var iToken Token = tokens[i]

				if isKeyword(iToken.Token) && iToken.Token != KW_PLUS {
					break
				} else {
					if iToken.Token == KW_PLUS {
						continue
					}

					// check equation
					VarInMemory := memory["var_"+iToken.Token]
					var isVar bool = false
					if VarInMemory != nil {
						_, err := strconv.Atoi(VarInMemory.(string))
						if err != nil {
							return
						}
						isVar = true
					}

					if iToken.tokenType == "number" || isVar {
						var equation []Token = []Token{}
						var result int
						for j := i; j < len(tokens); j++ {
							var jToken Token = tokens[j]
							if isKeyword(jToken.Token) && !isOperator(jToken.Token) {
								break
							}

							VarInMemory := memory["var_"+jToken.Token]
							if VarInMemory != nil {
								equation = append(equation, Token{VarInMemory.(string), jToken.Line, "number"})
							} else {
								equation = append(equation, jToken)
							}
						}

						result = solveEquation(equation)
						valueStr = strconv.Itoa(result)
						pos += len(equation)
						break
					} else {
						varInMemory := memory["var_"+iToken.Token]
						if varInMemory != nil {
							valueStr += varInMemory.(string)
						} else {
							valueStr += iToken.Token
						}
					}
				}

				// update pos
				posUpdateVar++
				i++
			}

			// check
			if isKeyword(varName.Token) {
				fmt.Println("Syntax error: variable name cannot be a keyword at line " + strconv.Itoa(varName.Line))
				return
			}
			if equalSign.Token != "=" {
				fmt.Println("Syntax error: expected '=' after variable name at line " + strconv.Itoa(equalSign.Line))
				return
			}
			if value.Token == ("EOF") {
				fmt.Println("Syntax error: expected a value after '=' at line " + strconv.Itoa(value.Line))
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

func solveEquation(equation []Token) int {
	// calculate the on hold
	var onHold []int = calculateOnHold(equation)
	var result int = 0

	for len(equation) > 1 {
		onHold = calculateOnHold(equation)
		var operatorIndex int = onHold[0]
		var operator Token = equation[operatorIndex]
		var leftVal Token = equation[operatorIndex-1]
		var rightVal Token = equation[operatorIndex+1]
		var iResult int
		leftValInt, Lerr := strconv.Atoi(leftVal.Token)
		RightValInt, Rerr := strconv.Atoi(rightVal.Token)

		if Lerr != nil || Rerr != nil && !isOperator(leftVal.Token) && isOperator(rightVal.Token) {
			fmt.Println("Syntax Error: Invalid number in equation")
			return 0
		}
		switch {
		case operator.Token == KW_MULT:
			iResult = (leftValInt * RightValInt)
		case operator.Token == KW_DIV:
			iResult = (leftValInt / RightValInt)
		case operator.Token == KW_PLUS:
			iResult = (leftValInt + RightValInt)
		case operator.Token == KW_MINUS:
			iResult = (leftValInt - RightValInt)
		}
		equation = slices.Replace(equation, operatorIndex-1, operatorIndex+2, Token{strconv.Itoa(iResult), leftVal.Line, "number"})
	}
	result, err := strconv.Atoi(equation[0].Token)
	if err != nil {
		fmt.Println("Error in parsing string to number!")
		return 0
	}
	return result
}

func calculateOnHold(equation []Token) []int {
	var onHold []int = []int{}
	for i := 0; i < len(equation); i++ {
		var iToken Token = equation[i]
		if !isOperator(iToken.Token) {
			continue
		}
		onHold = append(onHold, i)
	}

	sort.Slice(onHold, func(i, j int) bool {
		iOperatorForce := operatorForce(equation[onHold[i]].Token)
		jOperatorForce := operatorForce(equation[onHold[j]].Token)

		if iOperatorForce == jOperatorForce {
			return onHold[i] > onHold[j]
		}

		if iOperatorForce > jOperatorForce {
			return true
		} else {
			return false
		}
	})

	return onHold
}
