package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	F "fmt"
	U "lemin/utils"
	S "strings"
)

type s_Global struct {
	Ants      int
	StartRoom s_Rooms
	EndRoom   s_Rooms
	rooms     []s_Rooms
	links     []s_Links
}

type s_Rooms struct {
	r_name string
	r_x    int
	r_y    int
}

type s_Links struct {
	room1 string
	room2 string
}

func checkFile(arg string) bool {
	if len(arg) == 0 {
		return false
	}

	if S.Count(arg, ".txt") != 1 || arg[len(arg)-4:] != ".txt" {
		return false
	}

	if _, err := os.Stat(arg); err != nil {
		return false
	}

	return true
}

func ParseFile(file *os.File) (*s_Global, error) {

	var data s_Global
	var start, end int

	scanner := bufio.NewScanner(file)

	if U.IsEmptyFile(file) {
		return nil, errors.New("Empty File!")
	}

	// is number of ants
	if scanner.Scan() {
		if nbr, stat := U.IsOnlyInt(scanner.Text()); stat {
			data.Ants = nbr
		} else {
			return nil, errors.New("invalid data format, invalid number of aints")
		}
	}

	for scanner.Scan() {

		// standard checkLine  !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		if !U.IsValidLine(scanner.Text()) {
			return nil, errors.New("Invalid Data Format ")
		}
		
		//isComment #comment
		if U.IsComment(scanner.Text()) {
			continue
		}


		//isCommand ##start or ##end or ##err
		if S.HasPrefix(scanner.Text(), "##") {
			if scanner.Text() == "##start" {
				start++
				if scanner.Scan() && start == 1 && end == 0 && U.IsRoom(scanner.Text()) {
					parts := strings.Fields(scanner.Text())

					r_x, _ := strconv.Atoi(parts[1])
					r_y, _ := strconv.Atoi(parts[2])

					data.StartRoom = s_Rooms{parts[0], r_x, r_y}
					data.rooms = append(data.rooms, data.StartRoom)

				} else {
					return nil, errors.New("Invalid Data Format!")
				}

			} else if scanner.Text() == "##end" {
				end++
				if scanner.Scan() && start == 1 && end == 1 && U.IsRoom(scanner.Text()) {
					parts := strings.Fields(scanner.Text())

					r_x, _ := strconv.Atoi(parts[1])
					r_y, _ := strconv.Atoi(parts[2])

					data.EndRoom = s_Rooms{parts[0], r_x, r_y}
					data.rooms = append(data.rooms, data.EndRoom)

				} else { return nil, errors.New("Invalid Data Format!!") }
			}
			continue
		}

		//isRoom
		if U.IsRoom(scanner.Text()) {
			parts := strings.Fields(scanner.Text())
			r_x, _ := strconv.Atoi(parts[1])
			r_y, _ := strconv.Atoi(parts[2])
			data.rooms = append(data.rooms, s_Rooms{parts[0], r_x, r_y})
			continue
		}

		//islink between rooms
		if U.IsLink(scanner.Text()) {
			parts := strings.Split(scanner.Text(), "-")
			data.links = append(data.links, s_Links{parts[0], parts[1]})
			continue
		}
		// if !(comment || command || room || links)
		return nil, errors.New("Invalid Data Format Not [comment || command || room || links] ---> " + scanner.Text())
	}

	return &data, nil
}

func Printing(data *s_Global) {

	F.Println("number of ants-----> ", data.Ants)

	F.Println("Start and End Rooms-----> ", data.StartRoom, " ", data.EndRoom)

	F.Println("Rooms ----->")
	for _, val := range data.rooms {
		F.Println(val.r_name, " ", val.r_x, " ", val.r_y)
	}

	F.Println("Links ----->")
	for _, val := range data.links {
		F.Println(val.room1, "-", val.room2)
	}

}

func IsGlobalEmpty(g *s_Global) bool {
	if g.Ants == 0 {
		return true
	}
	if len(g.rooms) == 0 || len(g.links) == 0 {
		return true
	}
	if (g.StartRoom == s_Rooms{} || g.EndRoom == s_Rooms{}) {
		return true
	}

	return false
}

func IsDupRoomName(rooms []s_Rooms) bool {
	_map := make(map[string]bool)

	for _, room := range rooms {
		if _map[room.r_name] {
			return true
		}
		_map[room.r_name] = true
	}
	return false
}

func IsDupCoords(rooms []s_Rooms) bool {

	for i, room := range rooms {
		for j := i + 1; j < len(rooms); j++ {
			if room.r_x == rooms[j].r_x && room.r_y == rooms[j].r_y {
				return true
			}
		}
	}
	return false

}

func IsDupTunnels(links []s_Links) bool {
	_map := make(map[string]bool)

	for _, tunnel := range links {
		if _map[tunnel.room1+tunnel.room2] {
			return true
		}
		_map[tunnel.room1+tunnel.room2] = true
	}
	return false
}

func IsMatchedTunnel(links []s_Links) bool {

	for _, link := range links {
		if link.room1 == link.room2 {
			return true
		}
	}
	return false

}

func IsUnknownRoom(data *s_Global) bool {

	_map := make(map[string]bool)

	for _, tunnels := range data.links {
		_map[tunnels.room1] = false
		_map[tunnels.room2] = false
	}

	for _, room := range data.rooms {
		_map[room.r_name] = true
	}

	for _, checkmap := range _map {
		if checkmap == false {
			return true
		}
	}
	return false
}

func ProcessData(data *s_Global) error {

	if IsGlobalEmpty(data) {
		return errors.New("Not Enough Tokens { Aints || Commands || Rooms || Links }")
	}
	if IsDupRoomName(data.rooms) {
		return errors.New("Duplicated Rooms Name")
	}
	if IsDupCoords(data.rooms) {
		return errors.New("Duplicated Rooms Coordinates")
	}
	if IsDupTunnels(data.links) {
		return errors.New("Duplicated Tunnels")
	}
	if IsMatchedTunnel(data.links) {
		return errors.New("Matched Rooms Linked")
	}
	if IsUnknownRoom(data) {
		return errors.New("Unknown Room In Tunnel")
	}
	//No path exists between start and end
	// if IsPathFound(data.links){ return errors.New("No Path Found")}

	return nil
}

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		os.Stderr.WriteString("ERROR: Invalid Usage [ go run . file.txt ] !\n")
		return
	}

	if !(checkFile(args[0])) {
		os.Stderr.WriteString("ERROR: Invalid Argument [" + args[0] + "] !\n")
		return
	}

	file, _ := os.Open(args[0])
	data, err := ParseFile(file)
	if err != nil {
		os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
		return
	}

	if err = ProcessData(data); err != nil {
		os.Stderr.WriteString("ERROR: invalid data format " + err.Error() + "\n")
		return
	}

	Printing(data)

	defer file.Close()

}
