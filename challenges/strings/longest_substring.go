package strings

func LongestSubstringOfTwoChars(input string) string {
	if len(input) == 0 {
		return ""
	}

	count := make(map[byte]int)
	left := 0
	maxLen := 0
	start := 0

	for right := 0; right < len(input); right++ {
		count[input[right]]++

		for len(count) > 2 {
			count[input[left]]--
			if count[input[left]] == 0 {
				delete(count, input[left])
			}
			left++
		}

		if len(count) == 2 && right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}
	}
	return input[start : start+maxLen]
}
