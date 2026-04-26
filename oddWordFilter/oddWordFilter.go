package oddwordfilter

import "strings"

func Solution(sentence string) string {
	// TODO: implement the solution here
	words := strings.Fields(sentence)
	var result []rune

	for _, word := range words {
		if len(word)&1 == 1 {
			for i := 0; i < len(word); i += 2 {
				result = append(result, rune(word[i]))
			}
		}
	}

	left, right := 0, len(result)
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return string(result) // Placeholder return
}
