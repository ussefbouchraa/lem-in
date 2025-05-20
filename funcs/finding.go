package lemin

import (
	"fmt"
	"sort"
	"sync"
)

func BfsWithBlock(connections map[string][]string, start, end string, blocked map[string]bool) []string {
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]
		if last == end {
			return path
		}
		for _, nb := range connections[last] {
			if visited[nb] || blocked[nb] {
				continue
			}
			visited[nb] = true
			newPath := append(append([]string{}, path...), nb)
			queue = append(queue, newPath)
		}
	}
	return nil
}

type result struct{ path []string }

func FindAllDisjointPaths(connections map[string][]string, start, end string) [][]string {
	const workers = 8
	pathChan := make(chan result, workers)
	done := make(chan struct{})
	var allPaths [][]string
	var wg sync.WaitGroup

	// 8 goroutines
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			blocked := map[string]bool{}
			for {
				select {
				case <-done:
					return
				default:
					p := BfsWithBlock(connections, start, end, blocked)
					if p == nil {
						return
					}
					pathChan <- result{p}
					for _, room := range p {
						if room != start && room != end {
							blocked[room] = true
						}
					}
				}
			}
		}()
	}

	// remov goroutine
	go func() {
		wg.Wait()
		close(pathChan)
	}()

	// add all paths
	for res := range pathChan {
		allPaths = append(allPaths, res.path)
	}

	// Filter paths
	final := [][]string{}
	used := map[string]bool{}
	for _, p := range allPaths {
		ok := true
		for _, room := range p {
			if room != start && room != end && used[room] {
				ok = false
				break
			}
		}
		if ok {
			final = append(final, p)
			for _, room := range p {
				if room != start && room != end {
					used[room] = true
				}
			}
		}
	}
	close(done)
	return final
}

func ChooseOptimalPaths(paths [][]string, totalAnts int) [][]string {
	if len(paths) == 0 {
		return nil
	}
	sort.Slice(paths, func(i, j int) bool { return len(paths[i]) < len(paths[j]) })

	bestSet := paths[:1]
	minRounds := simulateRounds(bestSet, totalAnts)

	for k := 2; k <= len(paths); k++ {
		subset := paths[:k]
		rounds := simulateRounds(subset, totalAnts)
		if rounds < minRounds {
			minRounds = rounds
			bestSet = subset
		}
	}
	return bestSet
}

func simulateRounds(subset [][]string, totalAnts int) int {
	n := len(subset)
	antsPer := make([]int, n)

	for a := 0; a < totalAnts; a++ {
		best := 0
		bestScore := len(subset[0]) + antsPer[0]
		for i := 1; i < n; i++ {
			score := len(subset[i]) + antsPer[i]
			if score < bestScore {
				bestScore = score
				best = i
			}
		}
		antsPer[best]++
	}

	max := 0
	for i, count := range antsPer {
		if count == 0 {
			continue
		}
		rounds := len(subset[i]) + count - 2
		if rounds > max {
			max = rounds
		}
	}
	return max
}

type Ant struct {
	Id   int
	Path []string
	Idx  int
}

func SimulateAntsWalk(paths [][]string, antCount int) {
	type Ant struct {
		id   int
		path []string
		pos  int
	}
	pathCount := len(paths)
	antsPerPath := make([]int, pathCount)
	for i := 0; i < antCount; i++ {
		best := 0
		for j := 1; j < pathCount; j++ {
			if len(paths[j])+antsPerPath[j] < len(paths[best])+antsPerPath[best] {
				best = j
			}
		}
		antsPerPath[best]++
	}

	var ants []Ant
	id := 1
	for i, count := range antsPerPath {
		for j := 0; j < count; j++ {
			ants = append(ants, Ant{id, paths[i], 1})
			id++
		}
	}

	for moved := true; moved; {
		moved = false
		line := ""
		for i := range ants {
			if ants[i].pos < len(ants[i].path) {
				if line != "" {
					line += " "
				}
				line += fmt.Sprintf("L%d-%s", ants[i].id, ants[i].path[ants[i].pos])
				ants[i].pos++
				moved = true
			}
		}
		if moved {
			fmt.Println(line)
		}
	}
}

func FindPaths(g *Game) {
	all := FindAllDisjointPaths(g.Connections, g.StartRoom, g.EndRoom)
	best := ChooseOptimalPaths(all, g.AntCount)
	SimulateAntsWalk(best, g.AntCount)
}
