package Problem0013

var conversion = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0

	for i := range s {
		s1 := conversion[rune(s[i])]

		if i < len(s)-1 {
			s2 := conversion[rune(s[i+1])]

			if s1 < s2 {
				res -= s1
			} else {
				res += s1
			}
		} else {
			res += s1
		}
	}

	return res
}
