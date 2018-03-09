package main

import(
	"fmt"
)

var (
	// Endereco is a public variable
	Endereco string
	// Nome is a public variable
	Nome string

	comprou bool
)

func main() {
	fmt.Printf("Nome: %s\n", Nome)
	fmt.Printf("Endereco: %s\n", Endereco)
}