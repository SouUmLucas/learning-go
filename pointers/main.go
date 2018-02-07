package main

import "fmt"

type Imovel struct {
	nome    string
	comodos int
}

func main() {
	casa := Imovel{}
	casa.nome = "Primeiro nome"
	casa.comodos = 20

	fmt.Printf("Primeiros valores da struct %p %+v\n", &casa, casa)

	changeAttributes(&casa)
	fmt.Printf("Valores mudados da struct %p %+v\n", &casa, casa)
}

func changeAttributes(imovel *Imovel) {
	imovel.nome = "Segundo nome"
	imovel.comodos = 40
}