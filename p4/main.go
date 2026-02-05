package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Print("Enter number of elements to generate: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rng.Intn(1000) // Generate numbers between 0 and 999
	}

	fmt.Println("Original array:", arr)
	selectionSort(arr)
	fmt.Println("Sorted array:  ", arr)
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
