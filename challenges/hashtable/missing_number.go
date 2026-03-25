package hashtable

import "math"

func MissingNumber(numbers []int) int {
	set := make(map[int]struct{})
	maxNum := math.MinInt64
	minNum := math.MaxInt64

	for _, num := range numbers {
		set[num] = struct{}{}

		if num < minNum {
			minNum = num
		}

		if num > maxNum {
			maxNum = num
		}
	}

	for i := minNum; i < maxNum; i++ {
		if _, exist := set[i]; !exist {
			return i
		}
	}
	return -1
}
