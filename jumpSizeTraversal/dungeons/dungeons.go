package main

func Solution(dungeon []int, health int) int {
	n := len(dungeon)
	if n == 0 {
		return -1
	}

	bestJump := -1
	maxRemainingHealth := 0

	for jump := 1; jump <= n; jump++ {
		currentHealth := health
		pos := 0

		for pos < n {
			currentHealth -= dungeon[pos]

			if currentHealth <= 0 || pos+jump >= n {
				break
			}
			pos += jump
		}

		if currentHealth > maxRemainingHealth {
			maxRemainingHealth = currentHealth
			bestJump = jump
		}
	}

	return bestJump
}
