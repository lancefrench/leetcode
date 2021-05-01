package problem0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt32
	sort.Ints(nums)

	for i, _ := range nums {
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if abs(target-sum) < abs(diff) {
				diff = target - sum
			}
			if sum < target {
				low += 1
			} else {
				high -= 1
			}
			if diff == 0 {
				break
			}
		}
	}
	return target - diff
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
