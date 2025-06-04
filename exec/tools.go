package exec

import (
	"fmt"
	"sort"
	H "lemin/header"
)

func FormatOutput(farm *H.Farm) string {
	output := fmt.Sprintf("%d\n", farm.AntsCount)
	output += "##start\n"
	startRoom := farm.StartRoom
	output += fmt.Sprintf("%s %d %d\n", startRoom.R_name, startRoom.R_x, startRoom.R_y)

	var roomNames []string
	for name, room := range farm.Rooms {
		if room.R_name != farm.StartRoom.R_name && room.R_name != farm.EndRoom.R_name {
			roomNames = append(roomNames, name)
		}
	}
	sort.Strings(roomNames)

	for _, name := range roomNames {
		room := farm.Rooms[name]
		output += fmt.Sprintf("%s %d %d\n", room.R_name, room.R_x, room.R_y)
	}

	output += "##end\n"
	endRoom := farm.EndRoom
	output += fmt.Sprintf("%s %d %d\n", endRoom.R_name, endRoom.R_x, endRoom.R_y)

	for _, tunnel := range farm.Tunnels {
		output += fmt.Sprintf("%s-%s\n", tunnel.Room1, tunnel.Room2)
	}

	return output
}
