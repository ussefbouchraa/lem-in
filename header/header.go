package header



type Farm struct {
	Count     int
	AntsCount int
	Rooms     map[string]Room
	Tunnels   map[string]Tunnel
	StartRoom Room
	EndRoom   Room
}

type Room struct {
	R_name string
	R_x    int
	R_y    int
}

type Tunnel struct {
	Room1 string
	Room2 string
}

