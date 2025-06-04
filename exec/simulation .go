package exec

import (
	"fmt"
	"strings"
)
import H "lemin/header"


type Ant struct {
	ID       int
	PathIdx  int
	Position int
	Path     Path
	Finished bool
}

func SimulateAntMovement(farm *H.Farm, paths []Path) []string {
	if len(paths) == 0 {
		return nil
	}
	
	var moves []string
	var ants []*Ant
	
	distribution := OptimalAntDistribution(farm.AntsCount, paths)
	
	antID := 1
	for pathIdx, count := range distribution {
		for i := 0; i < count; i++ {
			ant := &Ant{
				ID:       antID,
				PathIdx:  pathIdx,
				Position: 0,
				Path:     paths[pathIdx],
				Finished: false,
			}
			ants = append(ants, ant)
			antID++
		}
	}
	
	finishedCount := 0
	for finishedCount < farm.AntsCount {
		var currentMoves []string
		roomOccupied := make(map[string]bool)
		
		for _, ant := range ants {
			if ant.Finished {
				continue
			}
			
			if ant.Position == 0 {
				if ant.Position+1 < len(ant.Path) {
					nextRoom := ant.Path[ant.Position+1]
					if !roomOccupied[nextRoom] || (nextRoom == farm.EndRoom.R_name && finishedCount == farm.AntsCount) {
						ant.Position++
						roomOccupied[nextRoom] = true
						
						if ant.Position == len(ant.Path)-1 {
							ant.Finished = true
							finishedCount++
						}
						
						moveStr := fmt.Sprintf("L%d-%s", ant.ID, nextRoom)
						currentMoves = append(currentMoves, moveStr)
					}
				}
			} else if ant.Position < len(ant.Path)-1 {
				nextRoom := ant.Path[ant.Position+1]
				if !roomOccupied[nextRoom] || nextRoom == farm.EndRoom.R_name {
					ant.Position++
					roomOccupied[nextRoom] = true
					
					if ant.Position == len(ant.Path)-1 {
						ant.Finished = true
						finishedCount++
					}
					
					moveStr := fmt.Sprintf("L%d-%s", ant.ID, nextRoom)
					currentMoves = append(currentMoves, moveStr)
				}
			}
		}
		
		if len(currentMoves) > 0 {
			moves = append(moves, strings.Join(currentMoves, " "))
		}
	}
	
	return moves
}

func OptimalAntDistribution(numAnts int, paths []Path) []int {
	if len(paths) == 0 {
		return nil
	}

	distribution := make([]int, len(paths))

	for i := 0; i < numAnts; i++ {
		bestPath := 0
		minFinishTime := len(paths[0]) - 1 + distribution[0]

		for j := 1; j < len(paths); j++ {
			finishTime := len(paths[j]) - 1 + distribution[j]
			if finishTime < minFinishTime {
				minFinishTime = finishTime
				bestPath = j
			}
		}
		distribution[bestPath]++
	}

	return distribution
}