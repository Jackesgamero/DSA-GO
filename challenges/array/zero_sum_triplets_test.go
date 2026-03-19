package array

import (
	"reflect"
	"sort"
	"testing"
)

/*
TestZeroSumTriplets tests solution(s) with the following signature and problem description:

	ZeroSumTriplets(list []int) [][]int

Given an array of numbers like {1, 2, -4, 6, 3} returns unique triplets from the numbers
with sum that equals zero like {-4, 1, 3} because -4+1+3=0.
*/

func ZeroSumTriplets(list []int) [][]int {
	sort.Ints(list)
	n := len(list)

	result := make([][]int, 0)

	var l, r, sum int
	for i := 0; i < n-2; i++ {
		//skip duplicates in i
		if i > 0 && list[i] == list[i-1] {
			continue
		}

		l, r = i+1, n-1
		for r > l {
			sum = list[i] + list[l] + list[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				result = append(result, []int{list[i], list[l], list[r]})

				// skip duplicates in l
				for l < r && list[l] == list[l+1] {
					l++
				}
				// skip duplicates in r
				for l < r && list[r] == list[r-1] {
					r--
				}

				l++
				r--
			}
		}
	}
	return result
}

func TestZeroSumTriplets(t *testing.T) {
	tests := []struct {
		list     []int
		triplets [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{}},
		{[]int{1, 2}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{}},
		{[]int{1, -4, 3}, [][]int{{-4, 1, 3}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{1, 2, -4, 6, 3}, [][]int{{-4, 1, 3}}},
		{[]int{-1, -2, -8, 6, 2, 3}, [][]int{{-8, 2, 6}, {-2, -1, 3}}},
		{[]int{1, -2, -4, 5, -2, 4, 1, 3}, [][]int{{-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}}},
	}

	for i, test := range tests {
		if got := ZeroSumTriplets(test.list); !reflect.DeepEqual(got, test.triplets) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.triplets, got)
		}
	}
}
