package main

import (
	"fmt"
)

func main() {
	userIds := []int{1, 2, 3, 2, 1, 5, 3, 1, 2, 1, 4, 5, 6}
	uniqueUserIds := processUserIds(userIds)
	fmt.Println(uniqueUserIds) // Expected output: [1 2 3 5 4 6]
}

func processUserIds(userIds []int) []int {
	userSet := make(map[int]struct{})

	// TODO: Add each userId to the userSet
	for _, id := range userIds {
		userSet[id] = struct{}{}
	}

	result := make([]int, 0, len(userSet))

	// TODO: Populate the result slice with unique IDs
	for id := range userSet {
		result = append(result, id)
	}

	return result
}
