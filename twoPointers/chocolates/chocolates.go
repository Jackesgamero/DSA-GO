package main

import (
	"fmt"
	"sort"
)

func findChocPairs(sweetness []int) [][]int {
	sort.Ints(sweetness)
	left, right := 0, len(sweetness)-1
	result := [][]int{}

	var sum int
	for left < right {
		sum = sweetness[left] + sweetness[right]

		if sum == 0 {
			result = append(result, []int{sweetness[right], sweetness[left]})
			left++
			right--
		} else if sum < 0 {
			left++
		} else {
			right--
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func main() {
	sweetness := []int{4, 3, -5, 5, -3, -4}
	result := findChocPairs(sweetness)
	for _, pair := range result {
		fmt.Printf("[%d, %d]\n", pair[0], pair[1])
	}
}
