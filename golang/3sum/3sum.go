package problem0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i-1] != nums[i] {
			twoSum(nums, i, &res)
		}
	}
	return res
}

func twoSum(nums []int, i int, res *[][]int) {
	seen := make(map[int]bool)
	j := i + 1
	for ; j < len(nums); j++ {
		complement := -nums[i] - nums[j]
		if seen[complement] {
			*res = append(*res, []int{nums[i], nums[j], complement})
			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j += 1
			}
		}
		seen[nums[j]] = true
	}
}
