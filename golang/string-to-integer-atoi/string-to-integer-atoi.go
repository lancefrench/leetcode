package problem0008

import "math"

func myAtoi(s string) int {
	sign := 1
	res := 0
	var i int

	if len(s) == 0 {
		return res
	}

	for i = 0; i < len(s); i++ {
		switch {
		case s[i] == ' ':
			continue
		case s[i] == '+':
			i++
		case s[i] == '-':
			sign = -1
			i++
		case s[i] < '0' || s[i] > '9':
			return 0
		}
		break
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		res = res*10 + int(s[i]-'0')
		if res > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			}
			return -math.MaxInt32 - 1
		}
	}
	return sign * res
}
