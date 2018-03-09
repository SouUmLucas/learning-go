package main

import "fmt"
import "runtime"
import "time"

func main() {
	currentOs := runtime.GOOS
	currentGoRoot := runtime.GOROOT

	fmt.Printf("Current Go root: %s", currentGoRoot)

	switch currentOs {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("*nix family")
	default:
		fmt.Println("Unsupported OS")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend...")
	default:
		fmt.Println("Time to go to work...")
	}
}