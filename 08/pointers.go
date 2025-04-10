// 08: Pointers

package main

import "fmt"

func increment(x *int) {
	*x = *x + 1
}

func main() {
	num := 5
	increment(&num)
	fmt.Println(num) // 6
}

/**
 * Pointers:
 * - Pointers store the memory address of a variable.
 * - &variable: Get the memory address of a variable.
 * - *pointer: Dereference a pointer to get the value it points to.
 */
