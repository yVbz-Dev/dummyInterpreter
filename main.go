package main

import (
	"fmt"
	"os"
)

type Token struct {
	Token     string
	Line      int
	tokenType string
}

type ASTNode struct {
	NodeAction string
	NodeArgs   map[string]any
}

var memory = make(map[string]any)

const (
	KW_PRINT               = "print"
	KW_VAR                 = "var"
	KW_PLUS                = "+"
	KW_MINUS               = "-"
	KW_MULT                = "*"
	KW_DIV                 = "/"
	KW_EQUAL               = "="
	KW_EQUAL_CONDITION     = "=="
	KW_NOT_EQUAL_CONDITION = "!="
	KW_IF                  = "if"
	KW_LEFT_PARENTHESIS    = "("
	KW_RIGHT_PARENTHESIS   = ")"
)

var keywords = map[string]bool{
	KW_PRINT:               true,
	KW_VAR:                 true,
	KW_PLUS:                true,
	KW_MINUS:               true,
	KW_MULT:                true,
	KW_DIV:                 true,
	KW_EQUAL:               true,
	KW_EQUAL_CONDITION:     true,
	KW_IF:                  true,
	KW_NOT_EQUAL_CONDITION: true,
	KW_LEFT_PARENTHESIS:    true,
	KW_RIGHT_PARENTHESIS:   true,
}

func main() {
	// get file
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <filename.e>")
		return
	}

	// read file
	filename := args[1]
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lexer(string(input))
}
