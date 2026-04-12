package main

import (
	"fmt"
	"sort"
)

type pair struct {
	value int
	index int
}

func FindAndReplace(X []int, Y []int) []int {
	X_sorted := make([]pair, len(X))
	Y_sorted := make([]pair, len(Y))

	for i, v := range X {
		X_sorted[i] = pair{value: v, index: i}
	}
	for i, v := range Y {
		Y_sorted[i] = pair{value: v, index: i}
	}

	sort.Slice(X_sorted, func(i, j int) bool {
		return X_sorted[i].value < X_sorted[j].value
	})
	sort.Slice(Y_sorted, func(i, j int) bool {
		return Y_sorted[i].value < Y_sorted[j].value
	})

	result := make([]int, len(Y))
	j := 0
	for i := 0; i < len(Y); i++ {
		target := float64(Y_sorted[i].value) / 2.0

		for j < len(X)-1 && float64(X_sorted[j+1].value) <= target {
			j++
		}

		best := X_sorted[j]

		if j < len(X)-1 {
			p1 := X_sorted[j]
			p2 := X_sorted[j+1]

			d1 := target - float64(p1.value)
			d2 := float64(p2.value) - target

			if d2 < d1 {
				best = p2
			} else if d1 == d2 {
				if p2.index < p1.index {
					best = p2
				}
			}
		}
		result[Y_sorted[i].index] = Y[best.index]
	}

	return result
}

func main() {
	X := []int{4, 12, 9, 20}
	Y := []int{10, 20, 40, 50}
	result := FindAndReplace(X, Y)
	fmt.Println(result) // Output: [10 40 50 50]
}
