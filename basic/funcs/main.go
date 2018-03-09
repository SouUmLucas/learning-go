package main

import "fmt"
import "github.com/SouUmLucas/funcs/math"

func main() {
  resSum := math.Calculate(math.Sum, 5, 6)
  fmt.Printf("Sum: %d\n", resSum)
  resSum, rest := math.Div(8,3)
  fmt.Printf("Div: %d; Rest: %d\n", resSum, rest)
}

