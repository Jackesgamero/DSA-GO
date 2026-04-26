package oddwordfilter

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Coding tasks are fun and required", "danfeasst"},
		{"etrnal", ""},
		{"awe-inspiring", "giisiea"},
		{"a a", "aa"},
		{"python coding", ""},
		{"Hello, world!", ""},
		{"Mastering Advanced Looping and Implementation", "dagioLgiesM"},
		{"Stay hungry, Stay foolish.", ",rnh"},
		{"! o G e h C o s i M P", "PMisoCheGo!"},
		{"abcdefghijklmnopqrstuvwxyz", ""},
	}

	for _, tt := range tests {
		got := Solution(tt.input)
		if got != tt.want {
			t.Errorf("Solution(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
