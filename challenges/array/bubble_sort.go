package array

func BubbleSort(input []int) {
	n := len(input)

	for n > 1 {
		swapped := false

		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		n--
	}
}
