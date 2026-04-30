package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input string
		want  []rune
	}{
		{"BCAAB", []rune{'B', 'A', 'A', 'B', 'C'}},
		{"AB", []rune{'A', 'B'}},
		{"A", []rune{'A'}},
		{"CBA", []rune{'B', 'A', 'C'}},
		{"AAA", []rune{'A', 'A', 'A'}},
		{"XYZ", []rune{'X', 'Y', 'Z'}},
		{"ABCDE", []rune{'A', 'C', 'B', 'D', 'E'}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Solution(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
