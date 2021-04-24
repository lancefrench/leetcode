package problem0003

func lengthOfLongestSubstring(s string) int {
	asciiSet := [128]int{}
	max := 0
	for left, right := 0, 0; right < len(s); right++ {
		if asciiSet[s[right]] > left {
			left = asciiSet[s[right]]
		}

		if right-left+1 > max {
			max = right - left + 1
		}

		asciiSet[s[right]] = right + 1
	}
	return max
}
