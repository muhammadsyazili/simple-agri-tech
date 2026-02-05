package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	if isPalindrome(input) {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
