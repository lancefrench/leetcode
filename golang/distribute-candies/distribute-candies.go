package problem0575

func distributeCandies(candies []int) int {
	kinds := make(map[int]int)
	for _, v := range candies {
		kinds[v] = kinds[v] + 1
	}
	if len(kinds) >= len(candies)/2 {
		return len(candies) / 2
	}
	return len(kinds)
}
