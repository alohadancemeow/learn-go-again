// 04: Control Flow

package main

import "fmt"

func main() {
	// If-else statement
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	// Switch statement
	grade := "B"

	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good!")
	default:
		fmt.Println("Needs improvement.")
	}
}
