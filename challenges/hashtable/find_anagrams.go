package hashtable

import "sort"

func FindAnagrams(words []string) [][]string {
	wordMap := make(map[string][]string)

	for _, word := range words {
		key := SortString(word)
		wordMap[key] = append(wordMap[key], word)
	}

	result := [][]string{}
	for _, array := range wordMap {
		if len(array) > 1 {
			result = append(result, array)
		}
	}
	return result
}

func SortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
