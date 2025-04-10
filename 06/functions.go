// 06: Functions

package main

import "fmt"

// Function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	result := add(10, 20)
	fmt.Println("Sum:", result)

	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}
