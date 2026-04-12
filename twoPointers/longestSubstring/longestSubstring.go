package main

import (
	"fmt"
)

func LongestSubstringWithKDistinctCharacters(s string, K int) int {
	i := 0
	best := 0
	count := make(map[byte]int)

	for j := 0; j < len(s); j++ {
		count[s[j]]++

		for len(count) > K {
			c := s[i]
			count[c]--
			if count[c] == 0 {
				delete(count, c)
			}
			i++
		}

		if j-i+1 > best {
			best = j - i + 1
		}
	}

	return best
}

func main() {
	fmt.Println(LongestSubstringWithKDistinctCharacters("acaabcc", 2)) // Expected output: 4
	fmt.Println(LongestSubstringWithKDistinctCharacters("abaccc", 2))  // Expected output: 4
	fmt.Println(LongestSubstringWithKDistinctCharacters("aa", 1))      // Expected output: 2
}
