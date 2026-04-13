package recursion

func ReverseDigits(n int) int {
	return reverseRecursiveLy(n, 0, digitsCount(n))
}

func reverseRecursiveLy(n, i, l int) int {
	if n < 10 {
		return n
	}
	return (n%10)*powerOf(10, l-i-1) + reverseRecursiveLy(n/10, i+1, l)
}

func powerOf(base, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	return res
}

func digitsCount(n int) int {
	if n < 10 {
		return 1
	}

	return 1 + digitsCount(n/10)
}
