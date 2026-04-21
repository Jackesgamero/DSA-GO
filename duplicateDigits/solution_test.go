package duplicatedigits

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{1234, 11223344},
		{1, 11},
		{22, 2222},
		{9876, 99887766},
		{10000, 1100000000},
		{0, 0},
		{3333, 33333333},
		{4444, 44444444},
		{5555, 55555555},
		{6666, 66666666},
	}

	for _, tt := range tests {
		got := Solution(tt.input)
		if got != tt.want {
			t.Errorf("Solution(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
