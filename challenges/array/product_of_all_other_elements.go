package array

func ProductOfAllOtherElements(list []int) []int {
	result := make([]int, len(list))
	if len(list) < 2 {
		return result
	}

	n := len(list)
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * list[i-1]
	}

	rightProd := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProd
		rightProd *= list[i]
	}

	return result
}
