package main

import "fmt"

func main() {
	if desc, status := clientStatus(2); status {
		fmt.Printf("Qual a situação do client? %v", desc)
	}
}

func clientStatus(meses int) (desc string, status bool) {
	if meses > 0 {
		status = true
		desc = "O cliente está devendo..."
		return
	}

	status = false
	desc = "O cliente está em dia"
	return
}