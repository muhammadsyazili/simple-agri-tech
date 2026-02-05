package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var k int
		fmt.Scan(&k)
		fmt.Println(getPolycarpSequenceElement(k))
	}
}

func getPolycarpSequenceElement(k int) int {
	count := 0
	curr := 1
	for {
		if curr%3 != 0 && curr%10 != 3 {
			count++
		}
		if count == k {
			return curr
		}
		curr++
	}
}
