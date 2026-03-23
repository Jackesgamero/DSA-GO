package array

func FindDuplicate(list []int) int {
	elements := make(map[int]bool)

	for _, v := range list {
		if elements[v] {
			return v
		}
		elements[v] = true
	}
	return -1
}
