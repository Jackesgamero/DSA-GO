package recursion

func ClimbingStairs(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	return ClimbingStairs(n-1) + ClimbingStairs(n-2)
}
