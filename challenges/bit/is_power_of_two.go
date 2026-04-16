package bit

func IsPowerOfTwo(input int) bool {
	if input <= 0 {
		return false
	}
	return (input & (input - 1)) == 0
}
