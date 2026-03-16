package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(ReverseAndFlipCase("Hello, World!")) // Expected Output: "!DLROw ,OLLEh"
	fmt.Println(ReverseAndFlipCase("hello"))         // Expected Output: "OLLEH"
	fmt.Println(ReverseAndFlipCase("HELLO"))         // Expected Output: "olleh"
}

func ReverseAndFlipCase(str string) string {
	stack := []rune{}

	for _, ch := range str {
		//Add characters to the stack with case transformation
		stack = append(stack, ch)
	}

	reversed := make([]rune, len(stack))
	// Construct the reversed slice
	for i, _ := range reversed {
		n := len(stack) - 1
		if unicode.IsUpper(stack[n]) {
			reversed[i] = unicode.ToLower(stack[n])
		} else {
			reversed[i] = unicode.ToUpper(stack[n])
		}
		stack = stack[:n]
	}

	return string(reversed)
}
