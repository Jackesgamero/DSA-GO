/*
Embark on a challenge to find the maximum value in a cyclically sorted array.
The array, originally sorted, has been shifted, requiring a keen analytical approach.
Your task is to identify the highest value using an efficient search strategy and return
that value. Carefully examine the shifts in the array order and apply your coding skills
to locate the maximum element. This exercise evaluates both logic and technique.
*/

package main

import (
	"fmt"
)

func main() {
	vaults := []int{9, 8, 7, 3, 2, 10}
	fmt.Println(FindMax(vaults)) // Output: 10

	vaults2 := []int{5, 4, 3, 1, 6}
	fmt.Println(FindMax(vaults2)) // Output: 6

	vaults3 := []int{3, 2, 1, 10, 9}
	fmt.Println(FindMax(vaults3)) // Output: 10
}

func FindMax(vaults []int) int {
	left, right := 0, len(vaults)-1
	// Implement iterative binary search
	for left < right {
		mid := left + (right-left)/2

		if vaults[mid] > vaults[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left < len(vaults)-1 {
		return vaults[left+1]
	}

	return vaults[0]
}
