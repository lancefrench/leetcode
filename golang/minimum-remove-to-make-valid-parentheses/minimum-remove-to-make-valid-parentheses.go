package problem1249

import "strings"

func minRemoveToMakeValid(s string) string {
	res := make([]string, len(s))

	stack := []int{}
	for i, r := range s {
		// convert our unicode rune back to an ascii character
		c := string(r)
		switch c {
		case "(":
			stack = append(stack, i)
			res[i] = c
		case ")":
			if len(stack) == 0 {
				res[i] = ""
				continue
			}
			stack = stack[:len(stack)-1]
			res[i] = c
		default:
			res[i] = c
		}
	}

	for _, idx := range stack {
		res[idx] = ""
	}

	return strings.Join(res, "")
}
