package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ints = regexp.MustCompile("[0-9]+")

func main() {
	raw, err := os.ReadFile("./05/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	in := string(raw)

	stackwise := readStackTops(moveBoxesStackwise(parseInput(in)))
	fmt.Printf("Part 1 (stackwise): Tops of the stacks: %s\n", stackwise)

	queuewise := readStackTops(moveBoxesQueuewise(parseInput(in)))
	fmt.Printf("Part 2 (queuewise): Tops of the stacks: %s\n", queuewise)
}

func parseInput(in string) (Dock, []Instruction) {
	split := strings.Split(in, "\n\n")
	diagram, instructions := split[0], split[1]

	return NewDock(diagram), NewInstructionSet(instructions)
}

func moveBoxesStackwise(dock Dock, instructions []Instruction) Dock {
	d := dock

	for _, instruction := range instructions {
		for i := 0; i < instruction.Count; i++ {
			d.Stacks[instruction.Destination-1].Push(d.Stacks[instruction.Origin-1].Pop())
		}
	}

	return d
}

func moveBoxesQueuewise(dock Dock, instructions []Instruction) Dock {
	d := dock

	for _, instruction := range instructions {
		d.Stacks[instruction.Destination-1].Append(d.Stacks[instruction.Origin-1].Take(instruction.Count))
	}

	return d
}

func readStackTops(dock Dock) string {
	tops := make([]rune, 0, len(dock.Stacks))
	for _, stack := range dock.Stacks {
		tops = append(tops, stack.Peek())
	}
	return string(tops)
}

func reverse(a []string) []string {
	reversed := make([]string, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		reversed = append(reversed, a[i])
	}
	return reversed
}

func group(a string) []string {
	grouped := make([]string, 0, len(a))
	for i := 0; i <= len(a); i = i + 4 {
		grouped = append(grouped, strings.TrimSpace(a[i:i+3]))
	}
	return grouped
}

/// Dock ----------------------------------------------------------------------

type Dock struct {
	Stacks []Stack
}

func NewDock(diagram string) Dock {
	layers := reverse(strings.Split(diagram, "\n"))

	stacks := make([]Stack, len(group(layers[0])))
	for i, layer := range layers {
		if i != 0 {
			for j, container := range group(layer) {
				if container == "" {
					continue
				}
				stacks[j].Push(rune(container[1]))
			}
		}
	}

	return Dock{Stacks: stacks}
}

/// Instruction ---------------------------------------------------------------

type Instruction struct {
	Count       int
	Origin      int
	Destination int
}

func NewInstructionSet(instructions string) []Instruction {
	split := strings.Split(strings.TrimSpace(instructions), "\n")

	ins := make([]Instruction, 0, len(split))
	for _, instruction := range split {
		values := []int{}

		nums := ints.FindAllString(instruction, -1)
		for _, num := range nums {
			value, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			values = append(values, value)
		}

		ins = append(ins, Instruction{
			Count:       values[0],
			Origin:      values[1],
			Destination: values[2],
		})
	}

	return ins
}

/// Stack ---------------------------------------------------------------------

type Stack struct {
	boxes []rune
}

func (s *Stack) Push(container rune) {
	s.boxes = append(s.boxes, container)
}

func (s *Stack) Pop() rune {
	container := s.boxes[len(s.boxes)-1]
	s.boxes = s.boxes[:len(s.boxes)-1]
	return container
}

func (s *Stack) Append(containers []rune) {
	s.boxes = append(s.boxes, containers...)
}

func (s *Stack) Take(count int) []rune {
	containers := s.boxes[len(s.boxes)-count:]
	s.boxes = s.boxes[:len(s.boxes)-count]
	return containers
}

func (s *Stack) Peek() rune {
	return s.boxes[len(s.boxes)-1]
}

func (s *Stack) Display() {
	for _, box := range s.boxes {
		fmt.Printf("%s", string(box))
	}
}
