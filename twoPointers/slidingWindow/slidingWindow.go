package main

import (
	"fmt"
)

func maximumSum(numbers []int, k int) []int {
	n := len(numbers)
	left := 0

	sum := 0
	for i := 0; i < k; i++ {
		sum += numbers[i]
	}

	maxSum := sum
	bestIndex := 0

	for right := k; right < n; right++ {
		sum += numbers[right]
		sum -= numbers[left]
		left++

		if sum > maxSum {
			maxSum = sum
			bestIndex = left
		}
	}

	return []int{maxSum, bestIndex}
}

func main() {
	result1 := maximumSum([]int{2, 1, 5, 1, 3, 2}, 3)
	fmt.Printf("[%d, %d]\n", result1[0], result1[1]) // Expected output: [9, 2]

	result2 := maximumSum([]int{2, 1, 5, 1, 3, 2}, 2)
	fmt.Printf("[%d, %d]\n", result2[0], result2[1]) // Expected output: [6, 1]
}
