package main

import (
	"dayseven/fs"
	"dayseven/parser"
	"fmt"
	"os"
	"strings"
)

const (
	Limit       = 100_000
	DiskSize    = 70_000_000
	NeededSpace = 30_000_000
)

func main() {
	raw, err := os.ReadFile("./puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dir := parser.Parse(strings.Split(strings.TrimSpace(string(raw)), "\n"))
	fmt.Println(dir)

	fmt.Printf("Part 1: Total size of directories with size < 100000: %d\n", smallTally(dir, Limit))

	needToFree := NeededSpace - (DiskSize - dir.Size())
	fmt.Printf(`Part 2 
---------------------------
Disk:              %d
Need:              %d
Used:              %d
Need to Free         %d
Optimal deletion:    %d
`, DiskSize, NeededSpace, dir.Size(), needToFree, optimalDeletion(dir, needToFree))
}

func smallTally(dir *fs.Dir, limit int) int {
	total := 0
	size := dir.Size()
	if size <= limit {
		total += size
	}
	for _, dir := range dir.Children {
		total += smallTally(dir, limit)
	}
	return total
}

func optimalDeletion(dir *fs.Dir, needToFree int) int {
	optimal := DiskSize
	size := dir.Size()
	if needToFree < size && size < optimal {
		optimal = size
	}
	for _, dir := range dir.Children {
		o := optimalDeletion(dir, needToFree)
		if o < optimal {
			optimal = o
		}
	}

	return optimal
}
