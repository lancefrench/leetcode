package problem0524

import "strings"

func findLongestWord(s string, d []string) string {
	res := ""
	for _, str := range d {
		if len(str) > len(res) || len(str) == len(res) && str < res {
			if isSubstr(s, str) {
				res = str
			}
		}
	}
	return res
}

func isSubstr(s, str string) bool {
	for _, c := range str {
		p := strings.Index(s, string(c))
		if p == -1 {
			return false
		}
		s = s[p+1:]
	}
	return true
}
