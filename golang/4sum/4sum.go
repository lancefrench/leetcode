package problem0018

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 0, 4)
}

func kSum(nums []int, target int, start int, k int) [][]int {
	res := make([][]int, 0)
	if start == len(nums) || nums[start]*k > target || target > nums[len(nums)-1]*k {
		return res
	}
	if k == 2 {
		return twoSum(nums, target, start)
	}
	for i := start; i < len(nums); i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, set := range kSum(nums, target-nums[i], i+1, k-1) {
				tmp := append([]int{nums[i]}, set...)
				res = append(res, tmp)
			}
		}
	}
	return res
}

func twoSum(nums []int, target int, start int) [][]int {
	res := make([][]int, 0)
	low, high := start, len(nums)-1
	for low < high {
		sum := nums[low] + nums[high]
		if sum < target || (low > start && nums[low] == nums[low-1]) {
			low += 1
		} else if sum > target || (high < len(nums)-1 && nums[high] == nums[high+1]) {
			high -= 1
		} else {
			res = append(res, []int{nums[low], nums[high]})
			low += 1
			high -= 1
		}
	}
	return res
}
