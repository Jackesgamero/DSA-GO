package main

import (
	"fmt"
)

func PartitionLabels(s string) []int {
	last := make([]int, 26)
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	var result []int
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}

		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	return result
}

func main() {
	lengths := PartitionLabels("abacdcd")
	for _, len := range lengths {
		fmt.Print(len, " ")
	}
	// Should print: 3 4
}
