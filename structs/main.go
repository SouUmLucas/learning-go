package main

import (
	"fmt"
)

type Imovel struct {
	X int
	Y int
	Nome string
	Valor float64
}

func main() {
	casa := Imovel{}
	fmt.Printf("Casa %+v\n", casa)

	ape := Imovel{ 20, 30, "Meu primeiro imovel", 180000.0 }
	fmt.Printf("Ape: %+v\n", ape)

	monsHome := Imovel {
		Y: 5,
		X: 80,
		Valor: 200000,
		Nome: "Casa da mae",
	}
	fmt.Printf("Casa da mãe: %+v\n", monsHome)
}
