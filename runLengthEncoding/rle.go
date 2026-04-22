package main

import (
	"strconv"
	"strings"
	"unicode"
)

func EncodeRLE(s string) string {
	// TODO: Implement the Run-Length Encoding (RLE) algorithm to encode the input alphanumeric string.
	var current rune
	var count int
	var result strings.Builder

	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			if c == current {
				count++
			} else {
				if current != 0 {
					result.WriteRune(current)
					result.WriteString(strconv.Itoa(count))
				}
				current = c
				count = 1
			}
		}
	}

	if current != 0 {
		result.WriteRune(current)
		result.WriteString(strconv.Itoa(count))
	}

	return result.String()
}
