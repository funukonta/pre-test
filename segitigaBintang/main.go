// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	input := 5

	// style biasa
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Printf("\n")

	// go 1.22 style
	for i := range input {
		for range i + 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
