package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Room struct represents a room with an ID
type Room struct {
	ID int
}

const (
	StartMarker = "##start"
	EndMarker   = "##end"
)

// Game struct to hold the state of the game
type Game struct {
	Rooms       map[int]Room
	Connections map[int][]int
	StartRoom   int
	EndRoom     int
	AntCount    int
}

// Parse function to load and parse the input file
func (g *Game) Parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("errrrrrror opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	var isStart, isEnd bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// fmt.Println("line---> ", line)

		// Check for start and end markers
		if line == StartMarker || line == EndMarker {
			// fmt.Println("isStart---> ", line)
			if line == StartMarker {
				isStart = true
			} else {
				isEnd = true
				fmt.Println("isEnd---> ", line)

			}
			continue
		}
		// Parse tunnels
		if strings.Contains(line, "-") {
			// fmt.Println("line+---> ", line)
			extract := strings.Split(line, "-")
			if len(extract) == 2 {
				from, _ := strconv.Atoi(extract[0])
				to, _ := strconv.Atoi(extract[1])
				g.Connections[from] = append(g.Connections[from], to)
				g.Connections[to] = append(g.Connections[to], from)
				// fmt.Printf("value-->%v ---key %v ", g.Connections[from], from)
				// fmt.Printf("key--> %v ---key %v ", g.Connections[to], to)
				// break
			}
		}
		// Parse rooms with IDs and store start/end
		if parts := strings.Fields(line); len(parts) == 3 {
			// fmt.Println("parts---> ", parts)
			// break
			id, _ := strconv.Atoi(parts[0])
			g.Rooms[id] = Room{ID: id}
			// valide start & End
			if isStart {
				// fmt.Println("sssss",id)
				g.StartRoom = id
				isStart = false
			} else if isEnd {
				g.EndRoom = id
				isEnd = false
				// fmt.Println("xxxxxx",g.EndRoom)

			}
		}
		// fmt.Println("Rooms---> ", g.Rooms)
	}

	// Check for any parsing errors

	return nil
}

func BfsWithBlock(connections map[int][]int, start, end int, blocked map[int]bool) []int {
	// fmt.Println("conne---> ", connections)
	queue := [][]int{{start}} // [[1]]
	// fmt.Println("s----> ", queue[0])

	visited := map[int]bool{start: true} // {1:true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		// fmt.Println("x----> ", queue)
		// break
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
	// fmt.Println(queue)
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
			// fmt.Println("ooooooooooooo")
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
		// fmt.Println("llllll")

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

// Main function to execute the logic
func (g *Game) FindPaths() {
	// fmt.Println("Rooms:", g.Rooms)
	fmt.Println("StartRoom ID:", g.StartRoom)
	// fmt.Println("EndRoom ID:", g.EndRoom)
	// fmt.Println("Connections:", g.Connections)

	paths := FindAllDisjointPaths(g.Connections, g.StartRoom, g.EndRoom)
	fmt.Println("Disjoint Paths:", paths)
}

func main() {
	// Initialize the game
	game := Game{
		Rooms:       make(map[int]Room),
		Connections: make(map[int][]int),
	}

	// Parse the input file
	err := game.Parse("filetest.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Find paths based on parsed data
	game.FindPaths()
}
