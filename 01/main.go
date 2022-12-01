package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in, err := os.ReadFile("./01/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	elves, err := foldCalories(string(in))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	max3, err := maxes(3, elves)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	sum3 := 0
	for _, max := range max3 {
		sum3 += max
	}

	fmt.Printf("Part 1: The elf carrying the most calories has %d calories.\n", max3[len(max3)-1])
	fmt.Printf("Part 2: The three elves carrying the most calories have a total of %d calories.\n", sum3)
}

// I used this solution for the first part of the puzzle, the second part
// made it clear that folding the list of calories each elf had into a
// slice of ints would better set up getting the answer to both parts
// of the puzzle.
func mostCalories(in string) (int, int, error) {
	maxCalElf, maxCal := 0, 0

	elf, cal := 0, 0
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			if cal > maxCal {
				maxCalElf, maxCal = elf, cal
			}

			cal = 0
			elf++
			continue
		}

		c, err := strconv.Atoi(line)
		if err != nil {
			return 0, 0, err
		}
		cal += c
	}

	return maxCalElf, maxCal, nil
}

// I developed this solution to better suit both parts of the puzzle when
// the second part was revealed. It takes the puzzle input and returns a
// slice of the calories held by each elf.
func foldCalories(in string) ([]int, error) {
	elves := []int{0}

	elf := 0
	for _, line := range strings.Split(strings.TrimSpace(in), "\n") {
		if line == "" {
			elf++
			elves = append(elves, 0)
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		elves[elf] += cal
	}

	return elves, nil
}

// maxes returns the n highest values in the given slice of ints.
func maxes(n int, cals []int) ([]int, error) {
	if n > len(cals) {
		return nil, fmt.Errorf("there cannot be %d max values in a slice of %d elements, bud", n, len(cals))
	}

	sort.Ints(cals)
	return cals[len(cals)-n:], nil
}
