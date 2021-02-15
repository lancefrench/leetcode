func isBipartite(graph [][]int) bool {
	// create map of vertex, color
	redBlack := map[int]int{}
	for i := 0; i < len(graph); i++ {
		if _, ok := redBlack[i]; ok { // skip the connected and visited graph
			continue
		}
		// new graph
		if !bipartile(graph, i, 1, redBlack) {
			return false
		}
	}
	return true
}

func bipartile(graph [][]int, node int, color int, redBlack map[int]int) bool {
	if nd, ok := redBlack[node]; ok { // visited
		return nd == color // return false if its different color than it suppose to be.i.e same as parent color or return true.
	}
	// set color on the node and mark as visited.
	redBlack[node] = color
	// adjencey list.
	for i := 0; i < len(graph[node]); i++ {
		if !bipartile(graph, graph[node][i], -color, redBlack) {
			return false
		}
	}
	return true

}