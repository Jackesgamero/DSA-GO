package array

func RotateKSteps(list []int, k int) {
	n := len(list)
	if n < 2 || k%n == 0 {
		return
	}

	curr := list[0]
	index := 0
	for range n {
		pos := (index + k) % n
		list[pos], curr = curr, list[pos]
		index = pos
	}
}
