package main

func lexer(sourceCode string) {
	// tokenize!
	var tokens []Token = tokenize(sourceCode)
	var nodes []ASTNode = parser(tokens)
	runner(nodes)
}

func tokenize(line string) []Token {
	// vars
	var currToken string
	var currTokenLine int
	var currTokenType string
	var readingToken bool = false
	var lineNum int = 0
	var tokens = []Token{}

	// read trought the line!
	for i := 0; i < len(line); i++ {
		char := line[i]

		if char == '\n' {
			lineNum++
		}

		// cases
		switch {
		case char == '\n':
			if readingToken {
				tokens = append(tokens, Token{currToken, lineNum, currTokenType})
				currTokenType = ""
				currToken = ""
				currTokenLine = lineNum
				readingToken = false
			}
		case char == ' ':
			if readingToken && currTokenType != "string" {
				if char == '}' {
					currToken += string(char)
				}
				tokens = append(tokens, Token{currToken, lineNum, currTokenType})
				currTokenType = ""
				currToken = ""
				currTokenLine = lineNum
				readingToken = false
			} else if readingToken {
				currToken += string(char)
			}
		case char == '"':
			if readingToken {
				readingToken = false
				tokens = append(tokens, Token{currToken, lineNum, currTokenType})
				currToken = ""
				currTokenType = ""
				currTokenLine = lineNum
			} else {
				readingToken = true
				currTokenType = "string"
				currTokenLine = lineNum
			}
		case (char >= '0' && char <= '9') && currTokenType != "string":
			readingToken = true
			currTokenType = "number"
			currToken += string(char)
			currTokenLine = lineNum
		case char == '/' && line[i+1] == '/' && !readingToken:
			return tokens
		case isKeyword(string(char)) && currTokenType == "number":
			tokens = append(tokens, Token{currToken, lineNum, currTokenType})
			tokens = append(tokens, Token{string(char), lineNum, "operator"})
			currTokenType = ""
			currToken = ""
			currTokenLine = lineNum
			readingToken = false
		case char == '(' || char == ')':
			if readingToken {
				tokens = append(tokens, Token{currToken, lineNum, currTokenType})
			}
			tokens = append(tokens, Token{string(char), lineNum, "condition"})
			currTokenType = ""
			currToken = ""
			currTokenLine = lineNum
			readingToken = false
		case isOperator(string(char)) && readingToken && currTokenType == "command":
			tokens = append(tokens, Token{currToken, lineNum, currTokenType})
			tokens = append(tokens, Token{string(char), lineNum, "operator"})
			currTokenType = ""
			currToken = ""
			currTokenLine = lineNum
			readingToken = false
		case currTokenLine != lineNum:
			tokens = append(tokens, Token{currToken, lineNum, currTokenType})
			currTokenType = ""
			currToken = ""
			currTokenLine = lineNum
			readingToken = false
		case readingToken:
			currToken += string(char)
		default:
			if i > 0 {
				var oldChar string = string(line[i-1])
				if oldChar != " " && oldChar != "\n" {
					currToken += oldChar
				}
			}
			currTokenLine = lineNum
			readingToken = true
			currTokenType = "command"
			currToken += string(char)
		}

		if i+1 >= len(line) {
			if currToken == "" || currToken == " " {
				continue
			}

			// append the fucking token
			tokens = append(tokens, Token{currToken, lineNum, currTokenType})
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
