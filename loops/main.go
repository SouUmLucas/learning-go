package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("#", i)
	}

	finalizou := false
	volta := 0
	for !finalizou {
		volta ++
		fmt.Println("Volta #", volta)
		if volta > 10 {
			finalizou = true
		}
	}

	frase := "Eu gosto de programar em Go"

	for index, letter := range frase {
		fmt.Printf("index: %d Letra %q\n", index, letter)
	}
}
