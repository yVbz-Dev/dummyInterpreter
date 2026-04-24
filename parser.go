package main

func parser(tokens []Token) {
	// var program []ASTNode
	var pos int = 0
	var program []ASTNode
	for pos < len(tokens) {
		token := tokens[pos]

		switch {
		case token.Token == "print":
			// check
			if nextToken(tokens, pos).Token == ("EOF") && nextToken(tokens, pos).tokenType != "string" {
				panic("Syntax error: expected a string after print command at line " + string(token.Line) + " column " + string(token.Column))
			}
			// create node
			var node ASTNode
			node.NodeAction = token.Token
			node.NodeArgs = map[string]any{
				"value": nextToken(tokens, pos).Token,
			}
			pos++
			program = append(program, node)
		}
		pos++
	}
	runner(program)
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
