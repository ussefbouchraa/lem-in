package lemin

import "sync"

func BfsWithBlock(connections map[int][]int, start, end int, blocked map[int]bool) []int {
	queue := [][]int{{start}} // [[1]]
	visited := map[int]bool{start: true} // {1:true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		// fmt.Println("path---->", path)
		last := path[len(path)-1]
		// fmt.Println("last--->", last)
		if last == end {
			return path
		}

		for _, neighbor := range connections[last] {
			// fmt.Println("xxx-->", neighbor)
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
	type result struct {
		path []int
	}

	const workers = 8 // you can adjust this based on your machine's capacity
	pathChan := make(chan result, workers)
	done := make(chan struct{})
	var allPaths [][]int
	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			blocked := map[int]bool{}
			for {
				select {
				case <-done:
					return
				default:
					path := BfsWithBlock(connections, start, end, blocked)
					if path == nil {
						return
					}
					pathChan <- result{path}
					for _, room := range path {
						if room != start && room != end {
							blocked[room] = true
						}
					}
				}
			}
		}()
	}

	// Wait for all goroutines to finish their work
	go func() {
		wg.Wait()
		close(pathChan) // Close channel after all workers finish
	}()

	// Collect all results from the channel
	for res := range pathChan {
		allPaths = append(allPaths, res.path)
	}

	// De-duplicate disjoint paths
	final := [][]int{}
	used := map[int]bool{}
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

	// Stop all goroutines gracefully
	close(done)

	return final
}
