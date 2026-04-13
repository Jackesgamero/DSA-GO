package main

import (
	"fmt"
)

func Solution(S string, Q []struct{ c1, c2 rune }) []int {
	// TODO: Implement the solution
	precalc := [26][26]int{}

	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			c1 := rune('a' + i)
			c2 := rune('a' + j)

			maxLen := 0
			curr := 0

			for _, char := range S {
				if char != c1 && char != c2 {
					curr++
					if curr > maxLen {
						maxLen = curr
					}
				} else {
					curr = 0
				}
				precalc[i][j] = maxLen
			}

		}
	}

	result := make([]int, len(Q))

	for i, q := range Q {
		c1 := q.c1 - 'a'
		c2 := q.c2 - 'a'
		result[i] = precalc[c1][c2]
	}

	return result
}

func main() {
	result1 := Solution("abcccacba", []struct{ c1, c2 rune }{{'a', 'b'}, {'b', 'c'}})
	fmt.Println(result1) // Example output: [3, 1]

	result2 := Solution("intelliaiassistant", []struct{ c1, c2 rune }{{'a', 'i'}, {'n', 't'}})
	fmt.Println(result2) // Example output: [5, 11]
}
