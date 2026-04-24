package digittoletter

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1-2-3-4-5", "a-b-c-d-e"},
		{"a-b-c", "1-2-3"},
		{"1-a-3-c-5", "a-1-c-3-e"},
		{"z-y-x-w-v", "26-25-24-23-22"},
		{"a-26-b-25-c-24", "1-z-2-y-3-x"},
		{"13-9-14-15", "m-i-n-o"},
		{"12-1-18-9-1", "l-a-r-i-a"},
		{"19-15-12-21-20-9-15-14", "s-o-l-u-t-i-o-n"},
		{"a-b-c-1-2-3-x-y-z-24-25-26", "1-2-3-a-b-c-24-25-26-x-y-z"},
		{"16-9-20-8-15-14-3-8-1-18-13-1", "p-i-t-h-o-n-c-h-a-r-m-a"},
	}

	for _, test := range tests {
		got := Solution(test.input)
		if got != test.want {
			t.Errorf("Solution(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
