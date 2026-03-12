/*
Imagine you're a code analyst tasked with deciphering two sets of reports.
Each report is an array of strings. Your mission is to identify the unique
strings from the second report that have an anagram match in the first report
and return the sum of their lengths as an integer. Remember, anagrams are words
formed by rearranging the letters of another. Now, if a string doesn't find an
anagram in the other report or isn't unique, move on. Whether the reports are
short or extensive, ensure your code efficiently calculates this cosmic string sum!
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to return a sorted string, used to identify anagram signatures
func SortCharacters(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// Function to find anagrams in report2 that are unique and have a matching anagram in report1
func FindAnagrams(report1, report2 []string) int {
	sortedWordsInReport1 := make(map[string]bool)
	// TODO: fill in sortedWordsInReport1
	for _, word := range report1 {
		sortedWordsInReport1[SortCharacters(word)] = true
	}

	anagramsMatched := make(map[string]bool)
	lengthSum := 0

	for _, word := range report2 {
		// TODO: sort the letters in the word
		sorted := SortCharacters(word)

		// TODO: check if word has an anagram match
		// TODO: add to the counter if the match is found for the first time
		if sortedWordsInReport1[sorted] {
			lengthSum += len(word)
			anagramsMatched[word] = true
		}
	}

	return lengthSum
}

func main() {
	report1 := []string{"cat", "dog", "tac", "god", "act"}
	report2 := []string{"tca", "ogd", "atc", "taco"}
	result := FindAnagrams(report1, report2)
	fmt.Println(result) // output: 9

	report3 := []string{"rat", "tar", "bat", "tab", "bats"}
	report4 := []string{"tra", "art", "abr"}
	result2 := FindAnagrams(report3, report4)
	fmt.Println(result2) // output: 6
}
