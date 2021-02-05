func findLHS(nums []int) int {
	sum := 0
	if len(nums) == 0 {
		return sum
	}

	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if val, ok := m[k+1]; ok {
			if v+val > sum {
				sum = v + val
			}
		}
	}
	return sum
}