package main

import (
	"fmt"
	F "lemin/funcs"
)

// Room struct represents a room with an ID

func main() {

	// Parse the input file
	g, err := F.Parse("q.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Find paths based on parsed data
	F.FindPaths(g)
}
