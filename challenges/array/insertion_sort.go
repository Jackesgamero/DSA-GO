package array

func InsertionSort(input []int) {
	for i := 1; i < len(input); i++ {
		key := input[i]
		j := i - 1

		for j >= 0 && input[j] > key {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = key
	}
}
