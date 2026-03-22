package strings

import (
	"strings"
	"testing"
)

var belowTwenty = []string{
	"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
	"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen",
	"Seventeen", "Nineteen",
}

var tens = []string{
	"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
}

var thousands = []string{"", "Thousand", "Million", "Billion"}

/*
TestReadNumberInEnglish tests solution(s) with the following signature and problem description:

	func NumberInEnglish(num int) string

Given n a positive integer smaller than a Trillion, return the number in English words.

For example given 0, return "Zero".
For example given 34, return "Thirty Four".
For example given 10, return "Ten".
For example given 900000000001, return "Nine Hundred Billion One".
*/

func NumberInEnglish(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := []string{}
	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			res = append([]string{helper(num%1000) + " " + thousands[i]}, res...)
		}
		num /= 1000
	}

	//Remove possible white spaces
	return strings.TrimSpace(strings.Join(res, " "))
}

func helper(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return belowTwenty[num]
	} else if num < 100 {
		if num%10 == 0 {
			return tens[num/10]
		}
		return tens[num/10] + " " + belowTwenty[num%10]
	} else {
		if num%100 == 0 {
			return belowTwenty[num/100] + " Hundred"
		}
		return belowTwenty[num/100] + " Hundred " + helper(num%100)
	}
}

func TestReadNumberInEnglish(t *testing.T) {
	tests := []struct {
		number  int
		english string
	}{
		{0, "Zero"},
		{1, "One"},
		{2, "Two"},
		{10, "Ten"},
		{34, "Thirty Four"},
		{123456789, "One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine"},
		{1000000000, "One Billion"},
		{100000000000, "One Hundred Billion"},
		{900000000001, "Nine Hundred Billion One"},
	}

	for i, test := range tests {
		got := NumberInEnglish(test.number)
		if got != test.english {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.english, got)
		}
	}
}
