/*
Imagine you are developing a smart library system.
You're tasked with evaluating a list of check-out events,
each represented by a unique identifier number for the library books.
Occasionally, a particular book might be checked out more frequently,
precisely more than n/4 times, where n is the total number of check-out events.
If such a book exists, your task is to print its identifier. Non-negative integers
represent the identifiers, but ignore the cases when an identifier is zero.
If no such frequent check-out book is found, the system should return -1.
*/

package main

import (
	"fmt"
)

func main() {
	checkOuts1 := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}
	frequentBook1 := FrequentBook(checkOuts1)
	fmt.Println("Frequent Book:", frequentBook1) // Expected output: 1

	checkOuts2 := []int{5, 0, 5, 0, 5, 0, 5, 0, 1, 1, 1, 1, 1}
	frequentBook2 := FrequentBook(checkOuts2)
	fmt.Println("Frequent Book:", frequentBook2) // Expected output: 5

	checkOuts3 := []int{3, 2, 2, 1, 3, 2, 3, 0, 0, 1, 4, 1}
	frequentBook3 := FrequentBook(checkOuts3)
	fmt.Println("Frequent Book:", frequentBook3) // Expected output: -1
}

func FrequentBook(checkOuts []int) int {
	countMap := make(map[int]int)
	frequentCheckOutThreshold := len(checkOuts) / 4

	// TODO: Implement the solution to return the frequently checked out book identifier
	for _, id := range checkOuts {
		countMap[id]++
		if countMap[id] > frequentCheckOutThreshold && id != 0 {
			return id
		}
	}
	return -1
}
