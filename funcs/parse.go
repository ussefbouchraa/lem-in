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
)

func Parse() {
	
	rooms := make(map[int]*Room)
	connections := map[int][]int{}
	file, _ := os.Open("filetest.txt")
	scanner := bufio.NewScanner(file)

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
		rooms[id] = &Room{Id: id}

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

	fmt.Println("Rooms:", rooms)
	fmt.Println("StartRoom ID:", StartRoom)
	fmt.Println("EndRoom ID:", EndRoom)

	for from, to := range connections {
		fmt.Printf("%v => %v", from, to)
	}

	Path := FindAllDisjointPaths(connections, StartRoom, EndRoom)

	fmt.Println("SHORT PATH : ", Path)
}
