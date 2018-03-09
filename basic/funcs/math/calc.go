package math

/*
  Calculate receives the calculation and the numbers
*/
func Calculate(f func(int, int) int, num1 int, num2 int) int {
	return f(num1, num2)
}
