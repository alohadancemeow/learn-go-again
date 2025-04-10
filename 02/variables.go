// 02: Variables

package main

import "fmt"

func main() {
	var message string = "Hello"
	count := 3 // short-hand, type inferred as int

	fmt.Println(message, count)
}

/**
Go can declare variables with var or use short-hand :=.
- var x type = value: Explicit type
- x := value: Implicit type (only inside functions)
*/
