package digittoletter

import (
	"strconv"
	"strings"
	"unicode"
)

/*
1 <-> a
2 <-> b
....
26 <-> z
*/

func Solution(s string) string {
	runes := []rune(s)
	var result strings.Builder
	var val int
	num := ""

	for _, r := range runes {
		if unicode.IsDigit(r) {
			num += string(r)
		} else {
			if num != "" {
				val, _ = strconv.Atoi(num)
				result.WriteRune(rune(val + 96))
				num = ""
			}

			if unicode.IsLetter(r) {
				val := int(r) - 96
				result.WriteString(strconv.Itoa(val))
			} else {
				result.WriteRune(r)
			}
		}
	}

	if num != "" {
		val, _ = strconv.Atoi(num)
		char := rune(val + 96)
		result.WriteRune(char)
	}

	return result.String()
}
