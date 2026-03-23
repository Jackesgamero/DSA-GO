package strings

import "strings"

var belowTwenty = []string{
	"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
	"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen",
	"Seventeen", "Nineteen",
}

var tens = []string{
	"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
}

var thousands = []string{"", "Thousand", "Million", "Billion"}

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
