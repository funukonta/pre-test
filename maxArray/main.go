// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	arr := []int{3, 5, 1, 9, 2, 10}

	numberTerbesar := maxNumberInArray(arr)
	fmt.Println(numberTerbesar)
}

func maxNumberInArray(arr []int) int {
	var terbesar int
	for i := range arr {
		if terbesar < arr[i] {
			terbesar = arr[i]
		}
	}
	return terbesar
}
