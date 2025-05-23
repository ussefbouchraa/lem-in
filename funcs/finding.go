package lemin

import (
	"fmt"
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
				ok = false // check path
				break
			}
		}
		if ok {
			final = append(final, p) // append valid paths
			for _, room := range p { // aplly all rooms in this path ---> visitide
				if room != start && room != end {
					used[room] = true
				}
			}
		}
	}
	close(done)
	return final
}

func Check(paths [][]string) {
	fmt.Println("len how many paths have ", len(paths))
	for _, v := range paths {
		fmt.Println("P--> ", len(v))
	}
}

func Small(g int, paths [][]string) {
	for i := 0; i < len(paths); i++ {
		round := (len(paths[i]) - 1 + g - 1)
		fmt.Println("Round-----> ", round)

	}
	return
}

func ChoosingOptimalpath(paths [][]string, totalants int) [][]string {
	bestRounds := int(^uint(0) >> 1) // max int
	bestpaths := [][]string{}

	for k := 1; k <= len(paths); k++ { // try all combinations of paths

		selcted := paths[:k]
		maxlen := 0

		for _, P := range selcted {
			if len(P) > maxlen {
				maxlen = len(P)
			}
		}

		round := maxlen - 1 + (totalants / k)

		if totalants%k != 0 {
			round++
		}

		if round < bestRounds {
			bestRounds = round
			bestpaths = selcted
		}
		fmt.Println("----> ", round)

	}
	return bestpaths
}

func FindPaths(g *Game) {
	all := FindAllDisjointPaths(g.Connections, g.StartRoom, g.EndRoom)
	// fmt.Println("all----> ", all)
	// Check(all)
	// Small(g.AntCount, all)

	fmt.Println(ChoosingOptimalpath(all, g.AntCount))
}
