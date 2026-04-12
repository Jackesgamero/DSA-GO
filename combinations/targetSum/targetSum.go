package main

import (
	"fmt"
)

func findPairSum(targetSum int, numbers []int) []int {
	seen := make(map[int]bool)

	for _, num := range numbers {
		complement := targetSum - num

		if seen[complement] {
			return []int{complement, num}
		}

		seen[num] = true
	}

	return []int{}
}

func main() {
	numbers := []int{2, 13, 4, 7, 5, 15}
	target := 9
	result := findPairSum(target, numbers)
	fmt.Println(result)
}
