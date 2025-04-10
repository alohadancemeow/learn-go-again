// 07: Data Structures

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	// Arrays
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[0]) // Access element

	// Slices (dynamic arrays)
	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println(slice)

	// Maps (key-value pairs)
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println(ages["Alice"])

	// Structs (custom data types)
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.Name, p.Age)

}
