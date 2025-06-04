package parse

import (
	"bufio"
	"errors"
	"fmt"
	H "lemin/header"
	"os"
	"strconv"
	"strings"
)

func CheckFile(arg string) bool {
	if len(arg) == 0 {
		return false
	}
	if strings.Count(arg, ".txt") != 1 || arg[len(arg)-4:] != ".txt" {
		return false
	}

	if _, err := os.Stat(arg); err != nil {
		return false
	}

	return true
}

func ParseFile(file *os.File) (*H.Farm, error) {
	data := H.Farm{
		Rooms:   make(map[string]H.Room),
		Tunnels: make(map[string]H.Tunnel),
	}

	scanner := bufio.NewScanner(file)

	if IsEmptyFile(file) {
		return nil, errors.New("Empty File!")
	}

	if scanner.Scan() {
		if nbr, stat := IsOnlyInt(scanner.Text()); stat {
			data.AntsCount = nbr
		} else {
			return nil, errors.New("Invalid Number of ants")
		}
	}

	var start, end int
	for scanner.Scan() {
		data.Count++
		// Read line and skip if empty
		line := scanner.Text()
		if line == "" {	continue}

		// standard checkLine  !
		if !IsValidLine(line) {
			return nil, errors.New("Invalid Line Format")
		}
		//isComment #comment
		if IsComment(line) {
			continue
		}
		//isCommand ##start or ##end or ##err
		if strings.HasPrefix(line, "##") {
			if line == "##start" {
				start++
				if scanner.Scan() && start == 1 && end == 0 && IsRoom(scanner.Text()) {
					parts := strings.Fields(scanner.Text())
					R_x, _ := strconv.Atoi(parts[1])
					R_y, _ := strconv.Atoi(parts[2])
					data.StartRoom = H.Room{parts[0], R_x, R_y}
					data.Rooms[parts[0]] = data.StartRoom
				} else {
					return nil, errors.New("Invalid Start Room Format")
				}
			} else if line == "##end" {
				end++
				if scanner.Scan() && start == 1 && end == 1 && IsRoom(scanner.Text()) {
					parts := strings.Fields(scanner.Text())
					R_x, _ := strconv.Atoi(parts[1])
					R_y, _ := strconv.Atoi(parts[2])
					data.EndRoom = H.Room{parts[0], R_x, R_y}
					data.Rooms[parts[0]] = data.EndRoom
				} else {
					return nil, errors.New("Invalid End Room Format")
				}
			}
			continue
		}
		//isRoom
		if IsRoom(line) {
			parts := strings.Fields(line)
			R_x, _ := strconv.Atoi(parts[1])
			R_y, _ := strconv.Atoi(parts[2])
			data.Rooms[parts[0]] = H.Room{parts[0], R_x, R_y}
			continue
		}

		if IsLink(line) {
			parts := strings.Split(line, "-")
			key := parts[0] + "-" + parts[1]
			data.Tunnels[key] = H.Tunnel{parts[0], parts[1]}
			continue
		}
		// if !(comment || command || room || links)
		return nil, errors.New("Invalid Data Format: " + line)
	}

	return &data, nil
}

func IsGlobalEmpty(g *H.Farm) bool {
	if g.AntsCount == 0 {
		return true
	}
	if len(g.Rooms) == 0 || len(g.Tunnels) == 0 {
		return true
	}
	if (g.StartRoom == H.Room{} || g.EndRoom == H.Room{}) {
		return true
	}
	return false
}

func IsDupCoords(rooms map[string]H.Room) bool {
	coordMap := make(map[string]bool)
	for _, room := range rooms {
		key := fmt.Sprintf("%d,%d", room.R_x, room.R_y)
		if coordMap[key] {
			return true
		}
		coordMap[key] = true
	}
	return false
}

func IsMatchedTunnel(tunnels map[string]H.Tunnel) bool {
	for _, link := range tunnels {
		if link.Room1 == link.Room2 {
			return true
		}
	}
	return false
}

func IsUnknownRoom(data *H.Farm) bool {
	for _, link := range data.Tunnels {
		if _, ok := data.Rooms[link.Room1]; !ok {
			return true
		}
		if _, ok := data.Rooms[link.Room2]; !ok {
			return true
		}
	}
	return false
}

func ProcessData(data *H.Farm) error {

	if IsGlobalEmpty(data) {
		return errors.New("Not Enough Tokens { Aints || Commands || Rooms || Tunnels }")
	}
	if IsDupCoords(data.Rooms) {
		return errors.New("Duplicated Rooms Coordinates")
	}

	if IsMatchedTunnel(data.Tunnels) {
		return errors.New("Matched Rooms Linked")
	}
	if IsUnknownRoom(data) {
		return errors.New("Unknown Room In Tunnel")
	}
	//No path exists between start and end
	// if IsPathFound(data.links){ return errors.New("No Path Found")}

	return nil
}
