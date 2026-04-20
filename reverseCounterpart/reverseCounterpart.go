package main

func Solution(numbers []int) [][]int {
	exists := make(map[int]bool)
	for _, num := range numbers {
		exists[num] = true
	}

	var result [][]int

	for _, num := range numbers {
		rev := reverseNumber(num)

		if exists[rev] {
			result = append(result, []int{num, rev})
		}
	}

	return result
}

func reverseNumber(n int) int {
	res := 0
	for n > 0 {
		remainder := n % 10
		res = (res * 10) + remainder
		n = n / 10
	}
	return res
}
