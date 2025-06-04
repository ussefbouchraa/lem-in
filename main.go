package main

import (
	"fmt"
	E "lemin/exec"
	P "lemin/parse"

	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		os.Stderr.WriteString("ERROR: Invalid Usage [ go run . file.txt ] !\n")
		return
	}

	if !(P.CheckFile(args[0])) {
		os.Stderr.WriteString("ERROR: Invalid Argument [" + args[0] + "] !\n")
		return
	}

	file, _ := os.Open(args[0])
	defer file.Close()

	data, err := P.ParseFile(file)
	if err != nil {
		os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
		return
	}

	if err = P.ProcessData(data); err != nil {
		os.Stderr.WriteString("ERROR: invalid data format " + err.Error() + "\n")
		return
	}



	paths, err := E.FindPaths(data)
	moves := E.SimulateAntMovement(data, paths)

	fmt.Print(E.FormatOutput(data))
	fmt.Println()
	for _, move := range moves {
		fmt.Println(move)
	}

}
