func kWeakestRows(mat [][]int, k int) []int {
	res := make([]int, 0)
	for j := 0; j < len(mat[0]); j++ { //iterate column by column
		for i := 0; i < len(mat); i++ {
			if j == 0 {
				mat[i] = append(mat[i], 0) //append 0 at the end to handle all 1s case
			}
			if mat[i][j] == 0 && (j == 0 || mat[i][j-1] != 0) { //find first 0 of the row
				res = append(res, i) //append into output
				if len(res) == k {
					return res
				}
			}
		}
	}

	return res
}