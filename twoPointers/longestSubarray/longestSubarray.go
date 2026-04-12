package main

import (
	"fmt"
)

func GetLongestSubarray(array []int, k int) []int {
	// TODO: implement function
	n := len(array)
	left := 0
	sum := 0

	bestLen := 0
	bestStart := -1

	for right := 0; right < n; right++ {
		sum += array[right]

		for sum > k && left <= right {
			sum -= array[left]
			left++
		}

		if sum == k {
			currentLen := right - left + 1
			if currentLen > bestLen {
				bestLen = currentLen
				bestStart = left
			}
		}
	}

	if bestStart == -1 {
		return []int{}
	}

	return array[bestStart : bestStart+bestLen]
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	k := 5
	result := GetLongestSubarray(array, k)
	fmt.Println(result) // Expected output: [2, 3]
}
