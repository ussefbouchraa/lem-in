package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Id int
}

var (
	StartRoom, EndRoom int
	istart, isEnd      bool
	antcount           int
)

func Parse() {
	r := make(map[int]*Room)
	connections := map[int][]int{}
	file, _ := os.Open("filetest.txt")
	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		firstline := scanner.Text()
		antcount, _ = strconv.Atoi(firstline)
		fmt.Println("antsssss---->", antcount)
	}

	for scanner.Scan() {
		fileline := scanner.Text()

		if len(fileline) == 0 {
			continue
		}
		if fileline == "##start" || fileline == "##end" {
			if fileline == "##start" {
				istart = true
				continue
			} else {
				isEnd = true
				continue
			}
		}

		// Tunnels
		if strings.Contains(fileline, "-") {
			extract := strings.Split(fileline, "-")
			if len(extract) == 2 {
				from, _ := strconv.Atoi(extract[0])
				to, _ := strconv.Atoi(extract[1])
				connections[from] = append(connections[from], to)
				connections[to] = append(connections[to], from)
				continue
			}
		}
		// Rooms
		extract := strings.Split(fileline, " ")
		if len(extract) != 3 {
			continue
		}

		id, _ := strconv.Atoi(extract[0])
		r[id] = &Room{Id: id}

		if istart || isEnd {
			if istart {
				StartRoom = id
				istart = false
			} else {
				EndRoom = id
				isEnd = false
			}
		}
	}

	fmt.Println("Rooms:", r)
	fmt.Println("StartRoom ID:", StartRoom)
	fmt.Println("EndRoom ID:", EndRoom)
	fmt.Println("con---->", connections)

	paths := FindAllDisjointPaths(connections, StartRoom, EndRoom)
	fmt.Println("SHORT PATH : ", paths)

	SimulateAnts(paths, antcount)
}

func SimulateAnts(paths [][]int, totalAnts int) {
	type Ant struct {
		id   int
		path []int
		pos  int
	}

	assigned := []Ant{}
	antID := 1

	// Step 1: Smart distribution
	for antID <= totalAnts {
		for _, path := range paths {
			if antID > totalAnts {
				break
			}
			assigned = append(assigned, Ant{id: antID, path: path, pos: 0})
			antID++
		}
	}

	// Step 2: Simulate step-by-step movement
	step := 0
	for {
		step++
		moved := false
		// fmt.Printf("Step %d:\n", step)
		for i := 0; i < len(assigned); i++ {
			ant := &assigned[i]
			if ant.pos < len(ant.path)-1 {
				ant.pos++
				room := ant.path[ant.pos]
				fmt.Printf("L%d-%d ", ant.id, room)
				moved = true
			}
		}
		fmt.Println()
		if !moved {
			break
		}
	}
}
