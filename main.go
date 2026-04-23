package main

import (
	"bufio"
	"fmt"
	"os"
)

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
