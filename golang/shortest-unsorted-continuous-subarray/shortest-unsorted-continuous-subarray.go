package problem0581

import "sort"

func findUnsortedSubarray(nums []int) int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)

	left := 0
	right := len(nums) - 1
	for left < len(nums) {
		if nums[left] == copyNums[left] {
			left++
		} else {
			break
		}
	}

	for left < right && right > 0 {
		if nums[right] == copyNums[right] {
			right--
		} else {
			break
		}
	}
	return right - left + 1
}
