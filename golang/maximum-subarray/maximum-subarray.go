package problem0053

import (
	"math"
)

func maxSubArray(nums []int) int {
	max_so_far, max_ending_here := (-math.MaxInt32 - 1), 0

	for i := 0; i < len(nums); i++ {
		max_ending_here = max_ending_here + nums[i]
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}

		if max_ending_here < 0 {
			max_ending_here = 0
		}
	}
	return max_so_far
}
