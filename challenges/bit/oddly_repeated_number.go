package bit

func OddlyRepeatedNumber(list []int) int {
	if len(list) == 0 {
		return -1
	}
	oddTimes := 0
	for _, element := range list {
		oddTimes ^= element
	}
	return oddTimes
}
