func shortestToChar(s string, c byte) []int {
	// The first instance of the character c in s;
	currIndex := strings.Index(s, string(c))
	// Previous instance of the character c in s;
	prevIndex := -100000

	i := 0
	ans := make([]int, len(s))

	for i < len(s) {
		// compare which one is closer the currentIndex or the previous Index.
		ans[i] = int(math.Min(float64(i-prevIndex), math.Abs(float64(currIndex-i))))

		// Whenever we reach the currentIndex we have to update the prevIndex and currentIndex.
		if i == currIndex && i < len(s) {
			prevIndex = currIndex
			currIndex = prevIndex + 1 + strings.Index(s[i+1:], string(c))
			fmt.Print(currIndex, " ")
		}
		i++
	}
	return ans
}