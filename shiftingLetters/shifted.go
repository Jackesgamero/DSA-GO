package main

import (
	//"encoding"
	"fmt"
	"unicode"
)

func main() {
	textToEncrypt := "Hello, Go World!"
	shiftValue := 3
	encryptedText := ""

	var shifted rune
	var tmp int

	// TODO: Implement a loop that goes through each character in the string
	for _, char := range textToEncrypt {

		// TODO: If the character is alphabetic, shift it by the given value
		// Make sure to maintain the case (upper or lower) of the original letter
		if unicode.IsLetter(char) {
			if unicode.IsLower(char) {
				tmp = (shiftValue + int(char-'a')) % 26
				shifted = rune(tmp + 'a')
			} else {
				tmp = (shiftValue + int(char-'A')) % 26
				shifted = rune(tmp + 'A')
			}
			encryptedText += string(shifted)
		} else {
			// TODO: If the character is not alphabetic, add it as is to the encrypted string
			encryptedText += string(char)
		}

	}

	// TODO: Print the encrypted text
	fmt.Println(encryptedText)
}
