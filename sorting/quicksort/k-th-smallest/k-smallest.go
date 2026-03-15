package main

import (
	"fmt"
	"math"
)

func findKRarest(scores []int, k int) int {
	if scores == nil || len(scores) < k {
		return math.MinInt32
	}
	return kRarest(scores, 0, len(scores)-1, k)
}

func kRarest(arr []int, l, r, k int) int {
	if k > 0 && k <= r-l+1 {
		//use partition to find pivot's position
		pos := partition(arr, l, r)

		//if pivot is in correct position, return the answer
		if pos-l == k-1 {
			return arr[pos]
		}

		if pos-l > k-1 { //call kRarest recursively if position is greater than k
			return kRarest(arr, l, pos-1, k)
		} else { //call kRarest recursively if position is less than k
			return kRarest(arr, pos+1, r, k-pos+l-1)
		}
	}
	return math.MaxInt32
}

func partition(arr []int, l, r int) int {
	pivot := arr[l]
	i := l

	for j := l + 1; j <= r; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[l] = arr[l], arr[i]
	return i
}

func main() {
	scores := []int{1, 3, 5, 7, 2, 1, 6}
	fmt.Println(findKRarest(scores, 4)) // It should print 3
}
