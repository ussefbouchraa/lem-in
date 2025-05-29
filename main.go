package main

import (
	"fmt"
	F "lemin/funcs"
	"time"
)

func main() {
	start := time.Now()
	g, err := F.Parse("filetest.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	F.FindPaths(*&g)

	fmt.Println("Execution time:", time.Since(start))

}
