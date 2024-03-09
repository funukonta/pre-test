// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	input := -1
	a := factor(input)
	fmt.Println(a)
}

func factor(x int) int {
	if x > 0 {
		return x * factor(x-1)
	} else {
		return 1
	}
}
