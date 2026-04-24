package main

import (
	"bufio"
	"fmt"
	"os"
)

type Token struct {
	Token     string
	Line      int
	Column    int
	tokenType string
}

type ASTNode struct {
	NodeAction string
	NodeArgs   map[string]any
}

var memory = make(map[string]any)

const (
	KW_PRINT = "print"
	KW_VAR   = "var"
)

var keywords = map[string]bool{
	KW_PRINT: true,
	KW_VAR:   true,
}

func main() {
	// read file
	conteudo, err := os.Open("codigo.e")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo .e: ", err)
	}
	defer conteudo.Close()

	// bufio!
	scanner := bufio.NewScanner(conteudo)
	lexer(*scanner)
}
