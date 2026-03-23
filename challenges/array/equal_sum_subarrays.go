package array

func EqualSubArrays(list []int) [][]int {
	totalSum := 0
	for i := 0; i < len(list); i++ {
		totalSum += list[i]
	}

	if totalSum%2 != 0 {
		return [][]int{}
	}

	sum := 0
	for i := 0; i < len(list)-1; i++ {
		sum += list[i]
		if sum == totalSum/2 {
			return [][]int{list[:i+1], list[i+1:]}
		}
	}
	return [][]int{}

}
