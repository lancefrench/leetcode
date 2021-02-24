package problem0240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix)-1, len(matrix[0])-1
	i, j := m, 0
	for i >= 0 && j <= n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}
