package main

import (
	"fmt"
	"os"
	"strings"
)

type Shape = string

const (
	Rock     Shape = "rock"
	Paper    Shape = "paper"
	Scissors Shape = "scissors"
)

type Strat = int

const (
	Win  Strat = 6
	Draw Strat = 3
	Lose Strat = 0
)

var (
	// Shape points when played
	throwPoints = map[Shape]int{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}
	// Throw symbols the other player uses
	otherThrows = map[byte]Shape{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
	}
	// The symbols I use in the first strat
	myThrows = map[byte]Shape{
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}
	// Strategy guide given in part 1.
	strat1 = map[string]int{
		"A X": Draw,
		"A Y": Win,
		"A Z": Lose,
		"B X": Lose,
		"B Y": Draw,
		"B Z": Win,
		"C X": Win,
		"C Y": Lose,
		"C Z": Draw,
	}
	// Strategy guide given in part 2.
	strat2 = map[string]int{
		"A X": Lose,
		"A Y": Draw,
		"A Z": Win,
		"B X": Lose,
		"B Y": Draw,
		"B Z": Win,
		"C X": Lose,
		"C Y": Draw,
		"C Z": Win,
	}
)

func main() {
	in, err := os.ReadFile("./02/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Part 1: Total score with the strategy guide 1 is %d.\n", part1Score(strings.TrimSpace(string(in))))
	fmt.Printf("Part 2: Total score with the strategy guide 2 is %d.\n", part2Score(strings.TrimSpace(string(in))))
}

func throws(round string) (other, mine byte) {
	return round[0], round[2]
}

func part1Score(in string) int {
	total := 0
	for _, round := range strings.Split(in, "\n") {
		_, myThrow := throws(round)
		total += throwPoints[myThrows[myThrow]] + strat1[round]
	}
	return total
}

func part2Score(in string) int {
	total := 0
	for _, round := range strings.Split(in, "\n") {
		total += throwPoints[pickShape(round, strat2[round])] + strat2[round]
	}
	return total
}

// Pick a shape to throw based on the win, lose, draw strategy and what the other player has thrown.
func pickShape(round string, strat Strat) Shape {
	theirs, _ := throws(round)
	switch strat {
	case Draw:
		return otherThrows[theirs]
	case Win:
		switch otherThrows[theirs] {
		case Rock:
			return Paper
		case Scissors:
			return Rock
		case Paper:
			return Scissors
		}
	case Lose:
		switch otherThrows[theirs] {
		case Rock:
			return Scissors
		case Scissors:
			return Paper
		case Paper:
			return Rock
		}
	}
	panic("Why don't Go switch statements recognize exhaustion across enums...")
}
