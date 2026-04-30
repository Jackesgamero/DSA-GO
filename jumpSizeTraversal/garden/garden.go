package main

func largestStep(garden []int, start int, direction int) int {
	// TODO: implement your solution here
	if len(garden) == 1 && start == 0 {
		return 1
	}

	flowerSet := make(map[int]struct{})
	for _, flower := range garden {
		flowerSet[flower] = struct{}{}
	}
	flowers := len(flowerSet)

	n := len(garden)
	best := -1
	for jump := 1; direction*jump+start >= 0 && direction*jump+start < n; jump++ {
		visited := make(map[int]struct{})
		visited[garden[start]] = struct{}{}
		pos := start

		for pos+jump*direction >= 0 && pos+jump*direction < n {
			pos += jump * direction
			visited[garden[pos]] = struct{}{}
		}

		if len(visited) == flowers && jump > best {
			best = jump
		}
	}
	return best
}
