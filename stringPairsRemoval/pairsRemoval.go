package main

func Solution(s string) []rune {
	var removed []rune
	current := s

	for len(current) > 0 {
		var nextRound []byte

		if len(current) == 1 {
			removed = append(removed, rune(current[0]))
			break
		}

		for i := 0; i < len(current); i += 2 {
			if i+1 == len(current) {
				nextRound = append(nextRound, current[i])
				continue
			}

			char1 := current[i]
			char2 := current[i+1]

			if char1 < char2 {
				removed = append(removed, rune(char1))
				nextRound = append(nextRound, char2)
			} else if char1 > char2 {
				removed = append(removed, rune(char2))
				nextRound = append(nextRound, char1)
			} else {
				removed = append(removed, rune(char1))
				nextRound = append(nextRound, char2)
			}
		}
		current = string(nextRound)
	}

	return removed
}
