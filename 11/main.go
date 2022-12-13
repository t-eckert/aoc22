package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var (
	ReInteger   = regexp.MustCompile("[0-9]+")
	RePostEq    = regexp.MustCompile("=.*$")
	ReOperation = regexp.MustCompile(`\+|\-|\*|\/`)
)

func main() {
	raw, err := os.ReadFile("./11/puzzle.input")
	if err != nil {
		panic(err.Error())
	}

	{
		monkeys := []*Monkey{}
		for _, in := range strings.Split(strings.TrimSpace(string(raw)), "\n\n") {
			monkeys = append(monkeys, NewMonkey(in))
		}

		for round := 1; round <= 20; round++ {
			playRoundPart1(monkeys)
		}

		inspected := []int{}
		for m, monkey := range monkeys {
			fmt.Printf("Monkey %d: %d\n", m, monkey.Inspected)
			inspected = append(inspected, monkey.Inspected)
		}

		slices.Sort(inspected)
		fmt.Printf("Part 1: Monkey business %d.\n", inspected[len(inspected)-2]*inspected[len(inspected)-1])
	}

	{
		monkeys := []*Monkey{}
		for _, in := range strings.Split(strings.TrimSpace(string(raw)), "\n\n") {
			monkeys = append(monkeys, NewMonkey(in))
		}

		for round := 1; round <= 10000; round++ {
			playRoundPart2(monkeys)
		}

		inspected := []int{}
		for m, monkey := range monkeys {
			fmt.Printf("Monkey %d: %d\n", m, monkey.Inspected)
			inspected = append(inspected, monkey.Inspected)
		}

		slices.Sort(inspected)
		fmt.Printf("Part 2: Monkey business %d.\n", inspected[len(inspected)-2]*inspected[len(inspected)-1])
	}
}

type Monkey struct {
	Items     []int
	Op        func(int) int
	Test      func(int) int
	Divisor   int
	Inspected int
}

func NewMonkey(in string) *Monkey {
	lines := strings.Split(in, "\n")

	// Get the items the monkey starts with
	items := []int{}
	for _, match := range ReInteger.FindAllString(lines[1], -1) {
		item, err := strconv.Atoi(match)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	// Get the operation on the worry level
	statement := RePostEq.FindString(lines[2])
	operation := ReOperation.FindString(statement)
	operand, operandErr := strconv.Atoi(ReInteger.FindString(statement))

	op := func(old int) int {
		var o int
		if operandErr == nil {
			o = operand
		} else {
			o = old
		}

		switch operation {
		case "+":
			return old + o
		case "-":
			return old - o
		case "*":
			return old * o
		case "/":
			return old / o
		}

		return old
	}

	// Get the test the monkey uses on the worry level
	divisor, err := strconv.Atoi(ReInteger.FindString(lines[3]))
	if err != nil {
		panic(err.Error())
	}
	truePass, err := strconv.Atoi(ReInteger.FindString(lines[4]))
	if err != nil {
		panic(err.Error())
	}
	falsePass, err := strconv.Atoi(ReInteger.FindString(lines[5]))
	if err != nil {
		panic(err.Error())
	}

	test := func(worry int) int {
		if worry%divisor == 0 {
			return truePass
		}
		return falsePass
	}

	// Seek forgiveness for the parsing I've done. Create the monkey.
	return &Monkey{
		Items:     items,
		Op:        op,
		Test:      test,
		Divisor:   divisor,
		Inspected: 0,
	}
}

func playRoundPart1(monkeys []*Monkey) {
	fmt.Println("Playing round")
	for _, monkey := range monkeys {
		thrown := 0
		for i := range monkey.Items {
			monkey.Inspected++

			idx := i - thrown
			item := monkey.Items[idx]

			worry := monkey.Op(item)
			worry = worry / 3
			monkey.Items[idx] = worry

			recipient := monkey.Test(worry)

			monkey.Items = append(monkey.Items[:idx], monkey.Items[idx+1:]...)
			monkeys[recipient].Items = append(monkeys[recipient].Items, worry)
			thrown++
		}
	}

	for m, monkey := range monkeys {
		fmt.Printf("Monkey %d: %+v\n", m, monkey.Items)
	}
	fmt.Println()
}

func playRoundPart2(monkeys []*Monkey) {
	// Sneaky trick because all of the divisors are prime
	p := 1
	for _, monkey := range monkeys {
		p *= monkey.Divisor
	}

	for _, monkey := range monkeys {
		thrown := 0
		for i := range monkey.Items {
			monkey.Inspected++

			idx := i - thrown
			item := monkey.Items[idx]

			worry := monkey.Op(item % p)
			monkey.Items[idx] = worry

			recipient := monkey.Test(worry)

			monkey.Items = append(monkey.Items[:idx], monkey.Items[idx+1:]...)
			monkeys[recipient].Items = append(monkeys[recipient].Items, worry)
			thrown++
		}
	}
}
