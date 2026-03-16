package main

import (
	"fmt"
)

func main() {
	fmt.Println(areBracketsBalanced("(){}[]")) // Output: true
	fmt.Println(areBracketsBalanced("([)]"))   // Output: false
	fmt.Println(areBracketsBalanced("{"))      // Output: false
}

func areBracketsBalanced(inputStr string) bool {
	bracketMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	openPar := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}

	var stack []rune

	// iterate over all characters in the input
	for _, char := range inputStr {
		if openPar[char] { // use a slice as a stack to store opening brackets found
			stack = append(stack, char)
		} else { // check if closing bracked has a matching pair in the stack
			if len(stack) == 0 || bracketMap[stack[len(stack)-1]] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	// return true only if all prackets match
	return len(stack) == 0
}
