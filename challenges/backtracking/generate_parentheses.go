package backtracking

func GenerateParentheses(n int) []string {
	combinations := []string{}
	backtrack(&combinations, "", 0, 0, n)
	return combinations
}

func backtrack(combinations *[]string, current string, open int, close int, max int) {
	if len(current) == max*2 {
		*combinations = append(*combinations, current)
		return
	}

	// Can I open a parentheses?
	if open < max {
		backtrack(combinations, current+"(", open+1, close, max)
	}

	// Can I close a parentheses?
	if close < open {
		backtrack(combinations, current+")", open, close+1, max)
	}
}
