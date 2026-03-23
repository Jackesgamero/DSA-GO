package strings

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
