package problem0056

import (
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	last := 0
	for i := 1; i < len(intervals); i++ {
		if res[last][1] >= intervals[i][0] {
			res[last][1] = max(res[last][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			last++
		}
	}
	return res
}
