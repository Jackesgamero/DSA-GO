package array

import "sort"

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
