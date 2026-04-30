package main

import "testing"

func TestTournament(t *testing.T) {
	tests := []struct {
		knights []int
		want    int
	}{
		{[]int{100, 50, 30, 20}, 3},
		{[]int{70, 80, 60}, 1},
		{[]int{30, 20, 10, 40, 50}, 2},
		{[]int{60, 70}, 1},
		{[]int{50, 20, 30, 40}, 1},
		{[]int{90, 10, 30}, 1},
		{[]int{30, 100}, 1},
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 1},
		{[]int{60, 40, 30, 20, 10}, 2},
		{[]int{100}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Tournament(tt.knights); got != tt.want {
				t.Errorf("Tournament() = %d, want %d", got, tt.want)
			}
		})
	}
}
