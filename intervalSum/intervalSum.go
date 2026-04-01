package main

import (
	"fmt"
)

func SumNumbers(queries [][2]int) []int64 {
	// TODO: Implement this method
	result := []int64{}
	for _, q := range queries {
		a, b := q[0], q[1]
		if a > b {
			a, b = b, a
		}

		sum := int64(b)*(int64(b)+1)/2 - int64(a)*(int64(a)-1)/2

		result = append(result, sum)
	}
	return result
}

func main() {
	queries := [][2]int{
		{1, 5},
		{5, 1},
		{1, 1},
		{500, 500},
		{1, 500},
		{123, 321},
	}

	results := SumNumbers(queries)
	for _, result := range results {
		fmt.Println(result)
	}
}
