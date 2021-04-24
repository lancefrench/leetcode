package problem0003

func lengthOfLongestSubstring(s string) int {
	asciiSet := [128]int{}
	max := 0
	for left, right := 0, 0; right < len(s); right++ {
		// if the last occurance of the character is greater than our left
		// window index, move the window forward so that we can never contain
		// a duplicate character
		if asciiSet[s[right]] > left {
			left = asciiSet[s[right]]
		}

		// Update max length if needed
		if right-left+1 > max {
			max = right - left + 1
		}

		// Set the value tracking our right index to one beyond
		// the last occurance of this character
		asciiSet[s[right]] = right + 1
	}
	return max
}
