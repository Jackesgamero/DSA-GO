package array

func AddSliceOfTwoNumbers(num1, num2 []int) []int {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0

	var result []int

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += num1[i]
			i--
		}

		if j >= 0 {
			sum += num2[j]
			j--
		}

		result = append([]int{sum % 10}, result...)
		carry = sum / 10
	}
	return result
}
