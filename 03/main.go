package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	in, err := os.ReadFile("./03/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 1: The priority sum is %d.\n", part1(strings.TrimSpace(string(in))))
	fmt.Printf("Part 2: The priority sum is %d.\n", part2(strings.TrimSpace(string(in))))
}

func part1(in string) int {
	sum := 0
	for _, line := range strings.Split(in, "\n") {
		sum += priority(inBoth(split(line)))
	}
	return sum
}

func part2(in string) int {
	size := 3
	group := make([]string, 3, 3)

	sum := 0
	for i, pack := range strings.Split(in, "\n") {
		group[i%size] = pack

		if i%size == size-1 {
			sum += priority(rune(common(group)[0][0]))
		}
	}

	return sum
}

func priority(b rune) int {
	if unicode.IsUpper(b) {
		return int(b) - 38
	}
	return int(b) - 96
}

func split(line string) (string, string) {
	return line[:len(line)/2], line[len(line)/2:]
}

func inBoth(a, b string) rune {
	for _, aByte := range a {
		for _, bByte := range b {
			if aByte == bByte {
				return aByte
			}
		}
	}
	return ' '
}

func common(packs []string) []string {
	if len(packs) < 2 {
		return packs
	}

	head, next := packs[0], packs[1]
	c := []rune{}
	for _, h := range head {
		for _, n := range next {
			if h == n {
				c = append(c, h)
			}
		}
	}

	return common(append([]string{string(c)}, packs[2:]...))
}
