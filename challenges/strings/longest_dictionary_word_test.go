package strings

import "testing"

/*
TestLongestDictionaryWordContainingKey tests solution(s) with the following signature and problem description:

	func LongestDictionaryWordContainingKey(key string, dic []string) string

Given a key as string, and a slice of strings containing a dictionary of words, return the longest
word that contains all letters of the key. Input will only contain lowercase letters a-z.

For example given "car" and {"rectify", "race", "archeology", "racoon"}, it should return "archeology",
because "archeology" is the longest word in the given set that contains "c","a",and "r".
*/

func LongestDictionaryWordContainingKey(key string, dic []string) string {
	var keyCount [26]int

	for i := 0; i < len(key); i++ {
		keyCount[key[i]-'a']++
	}

	longest := ""

	for _, word := range dic {
		var wordCount [26]int

		if len(word) <= len(longest) {
			continue
		}

		for i := 0; i < len(word); i++ {
			wordCount[word[i]-'a']++
		}

		valid := true
		for i := range 26 {
			if wordCount[i] < keyCount[i] {
				valid = false
				break
			}
		}

		if valid {
			longest = word
		}
	}
	return longest
}

func TestLongestDictionaryWordContainingKey(t *testing.T) {
	tests := []struct {
		input                           string
		dictionary                      []string
		longestWordContainingCharacters string
	}{
		{"a", []string{}, ""},
		{"a", []string{"c"}, ""},
		{"ab", []string{"cd"}, ""},
		{"ab", []string{"acd"}, ""},
		{"", []string{"abc"}, "abc"},
		{"a", []string{"a", "b", "c"}, "a"},
		{"a", []string{"a", "ba", "c", "cc"}, "ba"},
		{"a", []string{"a", "baa", "c"}, "baa"},
		{"abc", []string{"abc", "aabc", "aabbcc", "bbccaaccbbaa"}, "bbccaaccbbaa"},
		{"abc", []string{"abc", "abcdefghijklmn", "abcdefghijkl", "abcdef", "abcijkl"}, "abcdefghijklmn"},
		{"car", []string{"rectify", "race", "archeology", "racoon"}, "archeology"},
	}

	for i, test := range tests {
		got := LongestDictionaryWordContainingKey(test.input, test.dictionary)
		if got != test.longestWordContainingCharacters {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.longestWordContainingCharacters, got)
		}
	}
}
