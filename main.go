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
	NodeType string
	Value    any
	Children []ASTNode
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
