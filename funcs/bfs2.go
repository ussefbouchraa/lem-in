package lemin

func BfsWithBlock(connections map[int][]int, start, end int, blocked map[int]bool) []int {
	queue := [][]int{{start}}
	visited := map[int]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]

		if last == end {
			return path
		}

		for _, neighbor := range connections[last] {
			if visited[neighbor] || blocked[neighbor] {
				continue
			}
			visited[neighbor] = true
			newPath := append([]int{}, path...)
			newPath = append(newPath, neighbor)
			queue = append(queue, newPath)
		}
	}
	return nil
}

func FindAllDisjointPaths(connections map[int][]int, start, end int) [][]int {
	var Paths [][]int
	blocked := map[int]bool{}

	for {
		path := BfsWithBlock(connections, start, end, blocked)
		if path == nil {
			break
		}
		Paths = append(Paths, path)
		for _, room := range path {
			if room != start && room != end {
				blocked[room] = true
			}
		}
	}
	return Paths
}
