package duplicatedigits

func Solution(n int) int {
	if n == 0 {
		return 0
	}

	result := 0
	multiplier := 1

	for n > 0 {
		digit := n % 10
		duplicated := digit*10 + digit

		result += duplicated * multiplier

		multiplier *= 100
		n /= 10
	}

	return result
}
