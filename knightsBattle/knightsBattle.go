package main

func Tournament(knights []int) int {
	curr := make([]int, len(knights))
	copy(curr, knights)
	rounds := 0

	for {
		n := len(curr)
		if n <= 1 {
			break
		}

		allEqual := true
		for i := 1; i < n; i++ {
			if curr[i] != curr[0] {
				allEqual = false
				break
			}
		}
		if allEqual {
			break
		}

		var nextRound []int
		for i := 0; i < n; i++ {
			newStrength := curr[i] - curr[(i+1)%n]

			if newStrength > 0 {
				nextRound = append(nextRound, newStrength)
			}
		}
		curr = nextRound
		rounds++
	}
	return rounds
}
