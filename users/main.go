package main

import "fmt"
import "github.com/SouUmLucas/users/address"

var Username = "Lucas"

func main() {
	fmt.Printf("Welcome, %s\n", Username)
	fmt.Printf("You live in the city of %s, state of %s\n", address.Capital, address.State)
}
