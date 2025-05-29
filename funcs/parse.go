package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Room struct {
// 	ID string
// }

const (
	StartMarker = "##start"
	EndMarker   = "##end"
)

// Game struct to hold the state of the game
type Game struct {
	Rooms       []string
	Connections map[string][]string
	StartRoom   string
	EndRoom     string
	Ants        int
}

// Parse function to load and parse the input file
func Parse(filename string) (*Game, error) {

	g := &Game{
		Connections: make(map[string][]string),
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 1) read ant count
	if !scanner.Scan() {
		return nil, fmt.Errorf("empty file")
	}
	g.Ants, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))

	var isStart, isEnd bool
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		switch line {
		case StartMarker:
			isStart = true
			continue
		case EndMarker:
			isEnd = true
			continue
		}

		if strings.Contains(line, "-") { // tunnel
			parts := strings.Split(line, "-")
			g.Connections[parts[0]] = append(g.Connections[parts[0]], parts[1])
			g.Connections[parts[1]] = append(g.Connections[parts[1]], parts[0])
		} else if parts := strings.Fields(line); len(parts) == 3 { // room
			g.Rooms = append(g.Rooms, parts[0])
			if isStart {
				g.StartRoom = parts[0]
				isStart = false
			}
			if isEnd {
				g.EndRoom = parts[0]
				isEnd = false
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// fmt.Println("g.rooms-----> ", g.Rooms)
	return g, nil
}
