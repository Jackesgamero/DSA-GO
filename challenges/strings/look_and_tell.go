package strings

import "strings"

func LookAndTell(depth int) []string {
	if depth <= 0 {
		return []string{"-1"}
	}

	sequence := []string{"1"}

	for i := 1; i < depth; i++ {
		prev := sequence[i-1]
		var b strings.Builder
		count := 1

		for j := 0; j < len(prev); j++ {
			if j+1 == len(prev) || prev[j] != prev[j+1] {
				b.WriteByte(byte('0' + count))
				b.WriteByte(prev[j])
				count = 1
			} else {
				count++
			}
		}

		sequence = append(sequence, b.String())
	}

	return sequence
}
