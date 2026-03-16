package main

import (
	"fmt"
)

func findFirstFollowingSmallerElements(arr []int) []int {
	result := make([]int, len(arr))
	stack := []int{}

	// iterate over the collection
	for i := len(arr) - 1; i >= 0; i-- {
		// use the stack to find the first following smaller element
		for len(stack) != 0 && stack[len(stack)-1] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		// store output for each element of the collection
		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}

	return result
}

func main() {
	arr := []int{3, 7, 1, 7, 4, 3}
	result := findFirstFollowingSmallerElements(arr)
	fmt.Println(result) // Output: [1 1 -1 4 3 -1]

	// Additional tests
	arr2 := []int{4, 6, 2, 8, 1, 7}
	result2 := findFirstFollowingSmallerElements(arr2)
	fmt.Println(result2) // Output: [2 2 1 1 -1 -1]

	arr3 := []int{1, 1, 1, 1, 1}
	result3 := findFirstFollowingSmallerElements(arr3)
	fmt.Println(result3) // Output: [-1 -1 -1 -1 -1]
}
