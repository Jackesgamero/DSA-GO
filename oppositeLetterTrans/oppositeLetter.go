package main

import (
	"strings"
)

func Solution(s string) string {
	words := strings.Fields(s)
	n := len(words)
	if n == 0 {
		return ""
	}

	transformedWords := make([]string, n)
	for i, word := range words {
		var builder strings.Builder
		for _, r := range word {
			if r >= 'a' && r <= 'z' {
				// Opuesto minúscula: 'z' - (r - 'a')
				builder.WriteRune('z' - (r - 'a'))
			} else if r >= 'A' && r <= 'Z' {
				// Opuesto mayúscula: 'Z' - (r - 'A')
				builder.WriteRune('Z' - (r - 'A'))
			} else {
				builder.WriteRune(r)
			}
		}
		transformedWords[i] = builder.String()
	}

	var finalResult strings.Builder

	finalResult.WriteString(transformedWords[n-1])

	for i := 0; i < n-1; i++ {
		finalResult.WriteString(" ")
		finalResult.WriteString(transformedWords[i])
	}

	return finalResult.String()
}
