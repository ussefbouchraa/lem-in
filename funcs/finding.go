package lemin

import (
	"fmt"
)

var visitedTunnels = make(map[string]map[string]bool)

func BfsWithBlock(connections map[string][]string, start, end string) []string {
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		// fmt.Println("queue ", queue)
		path := queue[0]
		// fmt.Println("path ", path)

		queue = queue[1:] /////////////dequeue		///// FIFo \\\\\ OUT
		last := path[len(path)-1]
		// fmt.Println("last ", last)

		if last == end {
			// fmt.Println("++++ ", path)
			return path
		}

		for _, Nb := range connections[last] {

			if visitedTunnels[last] != nil && visitedTunnels[last][Nb] { // come back
				continue
			}

			if visited[Nb] {
				continue
			}

			visited[Nb] = true
			newPath := append(append([]string{}, path...), Nb)
			queue = append(queue, newPath) //////////// enquue
		}
	}
	return nil
}

func FindAllDisjointPaths(connections map[string][]string, start, end string) [][]string {
	var allPaths [][]string
	for {
		path := BfsWithBlock(connections, start, end)
		// fmt.Println(";;;;;;;;; ", path)
		if path == nil {
			break
		}
		allPaths = append(allPaths, path)

		for i, node := range path {
			if i+2 < len(path) && len(path) > 2 {
				if node != start && path[i+1] != end {
					if path[i+2] != end {

						if visitedTunnels[node] == nil {
							visitedTunnels[node] = make(map[string]bool)
						}
						visitedTunnels[node][path[i+1]] = true
						// fmt.Println("mappp------  ", visitedTunnels)
					}
				}
			} else if i+1 < len(path) && len(path) <= 2 {

				if visitedTunnels[node] == nil {
					visitedTunnels[node] = make(map[string]bool)
				}
				visitedTunnels[node][path[i+1]] = true

			}
		}
	}
	fmt.Println(connections)
	return allPaths
}

// func ChoosingOptimalpath(paths [][]string, totalants int) [][]string {

// 	bestRounds := int(^uint(0) >> 1) // max int
// 	bestpaths := [][]string{}

// 	for k := 1; k <= len(paths); k++ { // try all combinations of paths

// 		selcted := paths[:k]
// 		maxlen := 0

// 		for _, P := range selcted {
// 			if len(P) > maxlen {
// 				maxlen = len(P)
// 			}
// 		}

// 		round := maxlen - 1 + (totalants / k)

// 		if totalants%k != 0 {
// 			round++
// 		}

// 		if round < bestRounds {
// 			bestRounds = round
// 			bestpaths = selcted
// 		}

// 	}
// 	return bestpaths
// }

func FindPaths(g *Game) {
	all := FindAllDisjointPaths(g.Connections, g.StartRoom, g.EndRoom)
	for _, v := range all {
		fmt.Println("all[]-----", v)
		fmt.Println("len[v]-----", len(v))
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("lenall-----> ", len(all))
}
