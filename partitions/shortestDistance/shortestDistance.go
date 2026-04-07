package main

import (
	"fmt"
)

func CalculateMinDistances(wordList []string) map[string]int {
	result := make(map[string]int)
	lastOccurrence := make(map[string]int)

	for i, word := range wordList {
		if lastPos, found := lastOccurrence[word]; found {
			dist := i - lastPos

			if current, exists := result[word]; !exists || dist < current {
				result[word] = dist
			}
		}
		lastOccurrence[word] = i
	}

	return result
}

func main() {
	exampleInput := []string{"dog", "cat", "bird", "cat", "dog", "elephant", "dog"}
	result := CalculateMinDistances(exampleInput)
	// Expected output:
	// dog: 2
	// cat: 2
	// Note: The exact order of output pairs might differ.
	for k, v := range result {
		fmt.Printf("%s: %d\n", k, v)
	}
}
