package problem0413

func numberOfArithmeticSlices(A []int) int {
	res := 0
	if len(A) < 3 {
		return res
	}

	var i, j = 0, 0
	for i < len(A) {
		j = i + 2
		for j < len(A) && A[j]-A[j-1] == A[j-1]-A[j-2] {
			j++
		}
		j--
		res += (j - i - 1) * (j - i) / 2
		i = j
	}

	return res
}
