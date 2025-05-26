package lemin

import (
	"fmt"
	"sync"
)

var blocked = make(map[string]bool)
var mu sync.Mutex

func BfsWithBlock(connections map[string][]string, start, end string, blocked map[string]bool) []string {
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:] // dequeue
		last := path[len(path)-1]
		if last == end {
			mu.Lock()
			for _, v := range path {
				blocked[v] = true
			}
			mu.Unlock()
			return path
		}
		mu.Lock()
		for _, nb := range connections[last] {
			if visited[nb] || blocked[nb] {
				continue
			}
			visited[nb] = true
			newPath := append(append([]string{}, path...), nb)
			queue = append(queue, newPath) // enqueue
		}
		mu.Unlock()
	}
	return nil
}

type result struct{ path []string }

func FindAllDisjointPaths(connections map[string][]string, start, end string) [][]string {
	const workers = 8
	pathChan := make(chan result, workers) // buffred channel
	done := make(chan struct{})
	var allPaths [][]string
	var wg sync.WaitGroup
	// var mu sync.Mutex
	// 8 goroutines
	for w := 0; w < workers; w++ { // this loop is just runs 8 go routines
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:

					// blocked := make(map[string]bool) // private map of go routine
					p := BfsWithBlock(connections, start, end, blocked)
					if p == nil {
						return
					}

					pathChan <- result{p}
					// mu.Lock()
					// for _, room := range p {
					// 	if room != start && room != end {
					// 		blocked[room] = true
					// 	}
					// }
					// mu.Unlock()

				}
			}
		}()
	}

	
	go func() {
		wg.Wait()
		close(pathChan) 
	}()

	// add all paths
	for res := range pathChan {
		// fmt.Println(res.path)
		allPaths = append(allPaths, res.path)
	}

	final := [][]string{}
	used := map[string]bool{}
	for _, path := range allPaths {
		valid := true
		for _, room := range path {
			if room != start && room != end && used[room] {
				valid = false
				break
			}
		}
		if valid {
			final = append(final, path)
			for _, room := range path {
				if room != start && room != end {
					used[room] = true
				}
			}
		}
	}

	
	close(done)

	return final
}


func FindPaths(g *Game) {
	all := FindAllDisjointPaths(g.Connections, g.StartRoom, g.EndRoom)
	// for _, v := range all {
	// 	fmt.Println(len(v))
	// }
	fmt.Println("-----", all)

}
