// 05: Loops

package main

import "fmt"

func main() {
	// Standard for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Range loop
	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		fmt.Println(index, name)
	}
}

// Go only has the for loop. so sad!
