package main

import "testing"

func TestLargestStep(t *testing.T) {
	tests := []struct {
		garden    []int
		start     int
		direction int
		want      int
	}{
		{[]int{3, 1, 2, 1, 3, 2, 1}, 0, 1, 2},
		{[]int{1, 2, 3, 4, 5, 9, 2, 1, 3, 8, 2, 7, 1, 6}, 13, -1, 1},
		{[]int{1, 2, 3, 4, 5}, 0, 1, 1},
		{[]int{1}, 0, 1, 1},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 0, -1, -1},
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 9, -1, 1},
		{[]int{1, 2, 2, 4, 5, 5}, 0, 1, 1},
		{[]int{1, 1, 1, 1, 2, 1}, 0, 1, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, -1, -1},
		{[]int{1, 5, 2, 5, 3, 5, 4, 5}, 3, -1, -1},
	}

	for _, test := range tests {
		got := largestStep(test.garden, test.start, test.direction)
		if got != test.want {
			t.Errorf("largestStep(%v, %d, %d) = %d; want %d", test.garden, test.start, test.direction, got, test.want)
		}
	}
}
