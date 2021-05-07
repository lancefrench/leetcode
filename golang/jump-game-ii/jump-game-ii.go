package problem0045

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func jump(nums []int) int {
	jumps, currentJumpEnd, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = Max(farthest, i+nums[i])
		if i == currentJumpEnd {
			jumps++
			currentJumpEnd = farthest
		}
	}
	return jumps
}
