package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		dungeon []int
		health  int
		want    int
	}{
		{[]int{0, -1, 1, 0, -1}, 3, 1},
		{[]int{1, 0, -1, 1, 0}, 5, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 10, 10},
		{[]int{-3, -4, -2, -7, 8, -10, -3}, 14, 1},
		{[]int{1, 2, 3, 4, 5}, 20, 5},
		{[]int{100, 0, 0, 0, -10, 0, 0}, 110, 1},
		{[]int{0, -2, -4, -6, -8}, 10, 1},
	}

	for _, tt := range tests {
		got := Solution(tt.dungeon, tt.health)
		if got != tt.want {
			t.Errorf("Solution(%v, %d) = %d, want %d", tt.dungeon, tt.health, got, tt.want)
		}
	}
}
