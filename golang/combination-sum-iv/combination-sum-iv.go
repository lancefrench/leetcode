package problem0377

func combinationSum4(nums []int, target int) int {
	solutions := make([]int, target+1)

	solutions[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			// if
			if i-num < 0 {
				continue
			}

			solutions[i] += solutions[i-num]
		}
	}

	return solutions[target]
}
