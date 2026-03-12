package main

import (
	"fmt"
)

func FindNumber(array []int, targetNumber int) int {
	start := 0
	end := len(array) - 1

	for start <= end {
		// TODO: implement this
		mid := start + (end-start)/2

		if array[mid] == targetNumber {
			return mid
		} else if array[mid] > targetNumber {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1 // Number not found
}

func main() {
	sortedArray := []int{3, 7, 10, 20, 35, 40, 52, 60, 75}
	targetNumber := 20
	fmt.Println("Number found at index:", FindNumber(sortedArray, targetNumber))
}
