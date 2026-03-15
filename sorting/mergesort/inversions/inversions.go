package main

import (
	"fmt"
)

type Result struct {
	sorted     []int
	inversions int
}

func main() {
	numbers1 := []int{5, 4, 3, 2, 1}
	result1 := countInversions(numbers1)
	fmt.Println(result1.inversions) // Expected output: 10

	numbers2 := []int{-3, -2, -1, 0, 1}
	result2 := countInversions(numbers2)
	fmt.Println(result2.inversions) // Expected output: 0
}

func countInversions(arr []int) Result {
	if len(arr) <= 1 {
		return Result{sorted: arr, inversions: 0}
	}
	middle := len(arr) / 2
	left := countInversions(arr[:middle])
	right := countInversions(arr[middle:])
	result := mergeAndCountInversions(left.sorted, right.sorted)
	return Result{sorted: result.sorted, inversions: left.inversions + right.inversions + result.inversions}
}

func mergeAndCountInversions(left, right []int) Result {
	merged := make([]int, 0, len(left)+len(right))
	i, j, inversions := 0, 0, 0

	// TODO: Implement the remaining logic for merge and count inversions
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
			inversions += len(left) - i
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return Result{sorted: merged, inversions: inversions}
}
