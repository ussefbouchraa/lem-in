package exec

import (
	"fmt"
	H "lemin/header"
)

type Path []string

var visitedTunnels = make(map[string]map[string]bool)

func BfsWithBlock(connections map[string][]string, start, end string) []string {
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]

		queue = queue[1:]

		last := path[len(path)-1]

		if last == end {
			return path
		}

		for _, Nb := range connections[last] {

			if visited[Nb] || visitedTunnels[last] != nil && visitedTunnels[last][Nb] {
				continue
			}
			visited[Nb] = true
			newPath := append(append([]string{}, path...), Nb)
			queue = append(queue, newPath)
		}

	}
	return nil
}

func FindAllPathsMin(connections map[string][]string, start, end string) [][]string {
	var allPaths [][]string
	queue := [][]string{{start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]

		if last == end {
			allPaths = append(allPaths, path)
			continue
		}

		for _, nb := range connections[last] {
			if contains(path, nb) {
				continue // avoid reapet
			}
			newPath := append(append([]string{}, path...), nb)
			queue = append(queue, newPath)
		}
	}
	return allPaths
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func FindAllDisjointPaths(connections map[string][]string, start, end string) [][]string {
	var allPaths [][]string

	for {
		path := BfsWithBlock(connections, start, end)
		if path == nil {
			break
		}
		allPaths = append(allPaths, path)
		// sheck visited tunnel espect start ---> node & noode ---> End
		for i, node := range path {
			if i+1 < len(path) {

				if len(path) > 2 && node != start {

					if path[i+1] != end {
						if visitedTunnels[node] == nil {

							visitedTunnels[node] = make(map[string]bool)
						}
						visitedTunnels[node][path[i+1]] = true

					} else if len(path) == 3 {

						if visitedTunnels[node] == nil {
							visitedTunnels[node] = make(map[string]bool)
						}
						visitedTunnels[node][path[i+1]] = true
					}

				} else if len(path) == 2 {
					visitedTunnels[node] = make(map[string]bool)
					visitedTunnels[node][path[i+1]] = true
				}

			}
		}
	}
	return allPaths
}

func Decide(connections map[string][]string, start, end string, count int) [][]string {

	if count < 100 {
		paths := FindAllPathsMin(connections, start, end)
		return paths
	}

	paths := FindAllDisjointPaths(connections, start, end)

	return paths
}

func convertTunnelsToConnections(tunnels map[string]H.Tunnel) map[string][]string {
	connections := make(map[string][]string)
	for _, tunnel := range tunnels {
		connections[tunnel.Room1] = append(connections[tunnel.Room1], tunnel.Room2)
		connections[tunnel.Room2] = append(connections[tunnel.Room2], tunnel.Room1)
	}
	return connections
}

func FindPaths(g *H.Farm) ([]Path, error) {
	// Convert the Tunnels map to the expected connections format
	connections := convertTunnelsToConnections(g.Tunnels)

	// Use room names instead of Room structs
	rawPaths := Decide(connections, g.StartRoom.R_name, g.EndRoom.R_name, g.Count)

	if len(rawPaths) == 0 {
		return nil, fmt.Errorf("no path found from start to end")
	}

	// Convert [][]string to []Path
	var allPaths []Path
	for _, p := range rawPaths {
		allPaths = append(allPaths, Path(p))
	}

	// Now find the maximal set of disjoint paths
	disjointPaths := FindMaxDisjointPaths(allPaths)

	// Fallback if no disjoint set found (though unlikely)
	if len(disjointPaths) == 0 {
		disjointPaths = []Path{allPaths[0]}
	}
	return disjointPaths, nil
}

func findDisjointPathsSets(allPaths []Path, currentSet []Path, startIndex int, allSets *[][]Path) {
	if len(currentSet) > 0 {
		setCopy := make([]Path, len(currentSet))
		copy(setCopy, currentSet)
		*allSets = append(*allSets, setCopy)
	}
	for i := startIndex; i < len(allPaths); i++ {
		if canAddPathToSet(allPaths[i], currentSet) {
			currentSet = append(currentSet, allPaths[i])
			findDisjointPathsSets(allPaths, currentSet, i+1, allSets)
			currentSet = currentSet[:len(currentSet)-1]
		}
	}
}

func PathsShareIntermediate(path1, path2 Path) bool {
	for i := 1; i < len(path1); i++ {
		for j := 1; j < len(path2)-1; j++ {
			if path1[i] == path2[j] {
				return true
			}
		}
	}
	return false
}

func canAddPathToSet(newPath Path, pathSet []Path) bool {
	for _, existingPath := range pathSet {
		if PathsShareIntermediate(newPath, existingPath) {
			return false
		}
	}
	return true
}

func FindMaxDisjointPaths(allPaths []Path) []Path {
	if len(allPaths) == 0 {
		return nil
	}
	var allSets [][]Path
	findDisjointPathsSets(allPaths, []Path{}, 0, &allSets)

	if len(allSets) == 0 {
		return []Path{allPaths[0]}
	}

	bestSet := allSets[0]
	maxPaths := len(bestSet)
	bestTotalLength := calculateTotalLength(bestSet)

	for _, set := range allSets[1:] {
		if len(set) > maxPaths {
			bestSet = set
			maxPaths = len(set)
			bestTotalLength = calculateTotalLength(set)
		} else if len(set) == maxPaths {
			totalLength := calculateTotalLength(set)
			if totalLength < bestTotalLength {
				bestSet = set
				bestTotalLength = totalLength
			}
		}
	}
	return bestSet
}

func calculateTotalLength(paths []Path) int {
	total := 0
	for _, path := range paths {
		total += len(path)
	}
	return total
}
