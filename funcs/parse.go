package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
func Parse(filename string) (*Game, error) {
	
    g := &Game{
        Rooms:       make(map[int]Room),
        Connections: make(map[int][]int),
    }

    file, err := os.Open(filename)
    if err != nil { return nil, err }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // 1) read ant count
    if !scanner.Scan() { return nil, fmt.Errorf("empty file") }
    g.AntCount, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))

    var isStart, isEnd bool
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" { continue }

        switch line {
        case StartMarker:
            isStart = true; continue
        case EndMarker:
            isEnd = true;  continue
        }

        if strings.Contains(line, "-") {                    // tunnel
            parts := strings.Split(line, "-")
            from, _ := strconv.Atoi(parts[0])
            to,   _ := strconv.Atoi(parts[1])
            g.Connections[from] = append(g.Connections[from], to)
            g.Connections[to]   = append(g.Connections[to], from)
        } else if parts := strings.Fields(line); len(parts) == 3 { // room
            id, _ := strconv.Atoi(parts[0])
            g.Rooms[id] = Room{ID: id}
            if isStart { g.StartRoom = id; isStart = false }
            if isEnd   { g.EndRoom   = id; isEnd   = false }
        }
    }
    if err := scanner.Err(); err != nil { return nil, err }
    return g, nil
}
