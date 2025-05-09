package lemin

import "fmt"

func BFS(connections map[int][]int, startRoom, endRoom int) []int {
	queue := [][]int{{startRoom}} 
	fmt.Println("-----", queue)

	visited := make(map[int]bool)
	visited[startRoom] = true

	for len(queue) > 0 {
		path := queue[0]
		fmt.Println("-----path", path)

		queue = queue[1:]
		fmt.Println("-----queque", queue)

		current := path[len(path)-1]

		if current == endRoom {
			return path
		}

		for _, neighbor := range connections[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]int(nil), path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}
