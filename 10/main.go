package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./10/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	ins := parseInstructions(raw)
	sigs := compute(ins)

	total := 0
	for i := 19; i < len(sigs); i += 40 {
		total += sigs[i]
	}

	fmt.Printf("Part 1: Total signal strength is %d.\n", total)
	fmt.Println("Part 2: Rendering...")
	render(ins)
}

type Command = int

const (
	NOOP Command = 0
	ADDX Command = 1
)

type Instruction struct {
	Command Command
	Value   int
}

func parseInstructions(in []byte) []Instruction {
	ins := []Instruction{}
	for _, line := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		split := strings.Split(line, " ")

		var command int
		switch split[0] {
		case "noop":
			command = 0
		case "addx":
			command = 1
		}

		var value int
		if command == ADDX {
			value, _ = strconv.Atoi(split[1])
		}

		ins = append(ins, Instruction{command, value})
	}

	return ins
}

func compute(ins []Instruction) (signalStrengths []int) {
	X := 1
	C := 1

	for _, in := range ins {
		switch in.Command {
		case NOOP:
			signalStrengths = append(signalStrengths, X*C)
			C++
		case ADDX:
			signalStrengths = append(signalStrengths, X*C)
			C++
			signalStrengths = append(signalStrengths, X*C)
			X += in.Value
			C++
		}
	}

	return signalStrengths
}

func render(ins []Instruction) {
	X := 1
	C := 1

	for _, in := range ins {
		switch in.Command {
		case NOOP:
			draw(C, X)
			C++
		case ADDX:
			draw(C, X)
			C++
			draw(C, X)
			X += in.Value
			C++
		}
	}
}

func draw(C, X int) {
	col := (C - 1) % 40

	if X-2 < col && col < X+2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if col == 39 {
		fmt.Println()
	}
}
