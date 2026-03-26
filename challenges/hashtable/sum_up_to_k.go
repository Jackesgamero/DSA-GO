package hashtable

func SumUpToK(numbers []int, k int) []int {
	sums := make(map[int]int)

	for i, num := range numbers {
		if pos, ok := sums[num]; ok {
			return []int{pos, i}
		} else {
			sums[k-num] = i
		}
	}
	return []int{}
}
