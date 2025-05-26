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
		queue = queue[1:] // dequeue
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
			queue = append(queue, newPath) // enqueue
		}
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
	blocked := map[string]bool{} // private map of go routine
	var mu sync.Mutex
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
					mu.Lock()
					newblockedlist := (make(map[string]bool))

					for k, v := range blocked {
						newblockedlist[k] = v

					}
					mu.Unlock()
					p := BfsWithBlock(connections, start, end, newblockedlist)
					if p == nil {
						return
					}
					mu.Lock()
					for _, room := range p {
						if room != start && room != end && blocked[room] {
							mu.Unlock()
							return // skip: someone else took this path
						}
					}
					pathChan <- result{p}
					mu.Lock()
					for _, room := range p {
						if room != start && room != end {
							blocked[room] = true
						}
					}
					mu.Unlock()
				}
			}
		}()
	}

	// remov goroutine
	go func() {
		wg.Wait()
		close(pathChan) //  its job is only to wait until all 8 workers are done,
	}()

	// add all paths
	for res := range pathChan {
		fmt.Println(res.path)
		allPaths = append(allPaths, res.path)
	}

	close(done)
	return allPaths
}

// Main function to execute the logic
func FindPaths(g *Game) {
	all := FindAllDisjointPaths(g.Connections, g.StartRoom, g.EndRoom)
	fmt.Println("-----", all)

}
