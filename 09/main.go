package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./09/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	instructions := ParseInstructions(raw)
	fmt.Printf("Part 1: The tail visited %d locations.\n", play(instructions, 2))
	fmt.Printf("Part 2: The tail visited %d locations.\n", play(instructions, 10))
}

type Instruction struct {
	Direction string
	Delta     int
}

func ParseInstructions(raw []byte) []Instruction {
	instructions := make([]Instruction, 0, len(raw)/4)
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		split := strings.Split(line, " ")
		delta, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err.Error())
		}
		instructions = append(instructions, Instruction{
			Direction: split[0],
			Delta:     delta,
		})
	}
	return instructions
}

type Location struct {
	X int
	Y int
}

func play(instructions []Instruction, length int) int {
	// Locations the tail has been.
	p := map[Location]any{}

	H := Location{0, 0}
	rope := make([]Location, length-1)
	p[rope[len(rope)-1]] = nil

	for _, ins := range instructions {
		for step := 0; step < ins.Delta; step++ {
			// Move the head
			H = move(H, ins.Direction)
			for i, knot := range rope {
				if i == 0 {
					rope[i] = chase(H, knot)
				} else {
					rope[i] = chase(rope[i-1], knot)
				}
			}
			p[rope[len(rope)-1]] = nil
		}
	}

	return len(p)
}

func move(h Location, dir string) Location {
	switch dir {
	case "R":
		return Location{h.X + 1, h.Y}
	case "U":
		return Location{h.X, h.Y + 1}
	case "L":
		return Location{h.X - 1, h.Y}
	case "D":
		return Location{h.X, h.Y - 1}
	}

	return h
}

func chase(h, t Location) Location {
	dx := h.X - t.X
	dy := h.Y - t.Y

	if abs(dx) == 1 && abs(dy) == 1 {
		return t
	}

	if dx > 1 && dy == 0 {
		return Location{t.X + 1, t.Y}
	}
	if dx < -1 && dy == 0 {
		return Location{t.X - 1, t.Y}
	}
	if dy > 1 && dx == 0 {
		return Location{t.X, t.Y + 1}
	}
	if dy < -1 && dx == 0 {
		return Location{t.X, t.Y - 1}
	}

	if dx >= 1 && dy >= 1 {
		return Location{t.X + 1, t.Y + 1}
	}
	if dx >= 1 && dy <= -1 {
		return Location{t.X + 1, t.Y - 1}
	}
	if dx <= -1 && dy <= -1 {
		return Location{t.X - 1, t.Y - 1}
	}
	if dx <= -1 && dy >= 1 {
		return Location{t.X - 1, t.Y + 1}
	}

	return t
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
