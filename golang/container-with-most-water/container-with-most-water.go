package problem0011

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		x1, x2 := height[left], height[right]
		y := min(x1, x2)

		area := y * (right - left)
		if max < area {
			max = area
		}

		if x1 < x2 {
			left++
		} else {
			right--
		}
	}

	return max
}

// https://stackoverflow.com/questions/27516387/what-is-the-correct-way-to-find-the-min-between-two-integers-in-go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
