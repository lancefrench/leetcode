package problem0695

// A Depth-First Search (Iterative) approach
// TODO avoid searching squares we've already counted

func maxAreaOfIsland(grid [][]int) int {
	var res int
	for row := range grid {
		for col := range grid[row] {
			area := traverseIsland(grid, row, col)
			if res < area {
				res = area
			}
		}
	}
	return res
}

func traverseIsland(grid [][]int, row, col int) int {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == 0 {
		return 0
	}
	grid[row][col] = 0
	res := traverseIsland(grid, row+1, col)
	res += traverseIsland(grid, row-1, col)
	res += traverseIsland(grid, row, col+1)
	res += traverseIsland(grid, row, col-1)
	return res + 1
}
