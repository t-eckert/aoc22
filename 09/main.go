package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./09/test.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	instructions := ParseInstructions(raw)
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
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
