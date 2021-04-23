package problem0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		remainder := target - v
		// Check our map for the complement of our remainder
		if val, ok := m[remainder]; ok {
			return []int{val, k}
		}
		m[v] = k
	}
	return nil
}
