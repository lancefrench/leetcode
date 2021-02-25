package problem0856

func scoreOfParentheses(S string) int {
	stack, ans := []int{}, 0
	for _, c := range S {
		if c == '(' {
			stack = append(stack, ans)
			ans = 0
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = pop + max(ans*2, 1)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
