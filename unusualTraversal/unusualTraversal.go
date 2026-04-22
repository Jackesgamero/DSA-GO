package main

func UnusualTraversal(array []int) []int {
	// TODO: Implement array traversal starting from the middle, alternating between two elements from the left and right.
	var result []int

	n := len(array)
	mid := n / 2
	result = append(result, array[mid])
	left := mid - 2
	right := mid + 2

	for left >= 0 && right < n {
		result = append(result, array[left], array[left+1])
		result = append(result, array[right-1], array[right])
		left -= 2
		right += 2
	}

	if (n-1)%4 != 0 {
		result = append(result, array[0])
		result = append(result, array[n-1])
	}

	return result
}
