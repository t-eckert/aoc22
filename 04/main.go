package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, err := os.ReadFile("./04/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 1: %d assignments overlap completely.\n", countOverlapCompletely(strings.TrimSpace(string(in))))
	fmt.Printf("Part 2: %d assignments overlap at all.\n", countOverlapAtAll(strings.TrimSpace(string(in))))
}

type Assignment struct {
	From int
	To   int
}

func NewAssignment(assignment string) Assignment {
	split := strings.Split(assignment, "-")

	from, _ := strconv.Atoi(split[0])
	to, _ := strconv.Atoi(split[1])

	return Assignment{
		From: from,
		To:   to,
	}
}

func (a *Assignment) Contains(b Assignment) bool {
	return a.From <= b.From && b.To <= a.To
}

func (a *Assignment) Overlaps(b Assignment) bool {
	return a.To >= b.From && a.From <= b.To
}

func parseAssignments(line string) (Assignment, Assignment) {
	split := strings.Split(line, ",")
	return NewAssignment(split[0]), NewAssignment(split[1])
}

func countOverlapCompletely(in string) int {
	count := 0
	for _, line := range strings.Split(in, "\n") {
		a, b := parseAssignments(line)
		if a.Contains(b) || b.Contains(a) {
			count++
		}
	}
	return count
}

func countOverlapAtAll(in string) int {
	count := 0
	for _, line := range strings.Split(in, "\n") {
		a, b := parseAssignments(line)
		if a.Overlaps(b) {
			count++
		}
	}
	return count
}

// I used this for debugging.
func display(a, b Assignment) {
	fmt.Printf("A %d-%d\t:", a.From, a.To)
	for i := 1; i <= 99; i++ {
		if i >= a.From && i <= a.To {
			fmt.Print("*")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	fmt.Printf("B %d-%d\t:", b.From, b.To)
	for i := 1; i <= 99; i++ {
		if i >= b.From && i <= b.To {
			fmt.Print("*")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
