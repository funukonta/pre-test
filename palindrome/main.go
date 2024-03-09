// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	los := []string{
		"radar",
		"level",
		"devide",
		"civic",
		"evan",
	}

	for _, data := range los {
		isPalindrome(data)
	}

}

func isPalindrome(palindrome string) {
	rev_str := ""
	for _, str := range palindrome {
		rev_str = string(str) + rev_str
	}
	if rev_str == palindrome {
		fmt.Println(palindrome, true)
	} else {
		fmt.Println(palindrome, false)
	}
}
