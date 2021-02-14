type Node struct {
	x   int
	y   int
	dis int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	//handle special cases
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 && grid[0][0] == 0 {
		return 1
	}

	//define 8 connected directions
	directions := [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	//BFS queue definition
	queue := make([]Node, 0)

	//append start node
	node := Node{0, 0, 1}
	queue = append(queue, node)
	grid[0][0] = 1

	//process BFS
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			//iterate 8 directions
			for j := 0; j < len(directions); j++ {
				node := Node{}
				node.x = cur.x + directions[j][0]
				node.y = cur.y + directions[j][1]
				node.dis = cur.dis + 1

				//return the distance once reaches destination
				if node.x == n-1 && node.y == n-1 {
					return node.dis
				}

				//append valid connected node into queue
				if node.x >= 0 && node.x < n && node.y >= 0 && node.y < n && grid[node.x][node.y] == 0 {
					queue = append(queue, node)
					grid[node.x][node.y] = 1 //use grid value to indicate visited node
				}
			}
		}
	}
	return -1
}