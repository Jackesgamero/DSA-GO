package strings

import "testing"

/*
TestLongestSubstrings tests solution(s) with the following signature and problem description:

	func LongestSubstringOfTwoChars(input string) string

Given a string return the longest substring of two unique characters.

For example given "aabbc", return "aabb" because it's the longest substring that has two unique
characters "a" and "b".

Other substrings of "aabc" include:

* "aabbc", contains more than 2 unique characters.
* "bbc", shorter than "aabb".
*/

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

func TestLongestSubstrings(t *testing.T) {
	tests := []struct {
		input            string
		longestSubstring string
	}{
		{"", ""},
		{"a", ""}, // No occurrence of sequence of two characters
		{"ab", "ab"},
		{"abcc", "bcc"},
		{"aabbc", "aabb"},
		{"ada", "ada"},
		{"dafff", "afff"},
		{"abbdeeeddfffha", "deeedd"},
	}

	for i, test := range tests {
		got := LongestSubstringOfTwoChars(test.input)
		if got != test.longestSubstring {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.longestSubstring, got)
		}
	}
}
