package main

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

type Pos struct {
	X int
	Y int
}

func main() {
	raw, err := os.ReadFile("./12/puzzle.input")
	if err != nil {
		panic(err.Error())
	}

	terrain := parseTerrain(raw)
	domain := search(terrain, newVisited(terrain), [][]Pos{{start(terrain)}}, goal(terrain))
	fmt.Println(domain)
}

func parseTerrain(raw []byte) [][]int {
	terrain := [][]int{{}}

	row := 0
	for _, c := range raw[:len(raw)-1] {
		if c == '\n' {
			row++
			terrain = append(terrain, []int{})
		} else {
			terrain[row] = append(terrain[row], height(c))
		}
	}

	return terrain
}

func newVisited(terrain [][]int) [][]int {
	width, height := len(terrain[0]), len(terrain)

	visited := make([][]int, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]int, width)
	}

	return visited
}

func height(c byte) int {
	return int(c) - 97
}

func char(i int) byte {
	return byte(i + 97)
}

func posOf(terrain [][]int, value int) Pos {
	for y := range terrain {
		for x := range terrain[y] {
			if terrain[y][x] == value {
				return Pos{x, y}
			}
		}
	}
	return Pos{-1, -1}
}

func start(terrain [][]int) Pos {
	return posOf(terrain, -14)
}

func goal(terrain [][]int) Pos {
	return posOf(terrain, -28)
}

func contains(values []int, val int) bool {
	for _, value := range values {
		if value == val {
			return true
		}
	}

	return false
}

func search(terrain, visited [][]int, domain [][]Pos, goal Pos) [][]Pos {
	for _, pos := range domain[len(domain)-1] {
		visited[pos.Y][pos.X] = 1
	}

	if contains(subset(terrain, domain[len(domain)-1]), -28) {
		return domain
	}

	printTerrain(terrain, visited)

	next := next(terrain, visited, domain[len(domain)-1])
	return search(terrain, visited, append(domain, next), goal)
}

func next(terrain, visited [][]int, current []Pos) []Pos {
	next := []Pos{}

	return next
}

func canGo(terrain [][]int, from, to Pos) bool {
	fromHeight := terrain[from.Y][from.X]
	toHeight := terrain[to.Y][to.X]

	// Start and end cases
	if fromHeight == -14 && toHeight < 2 {
		return true
	}
	if toHeight == -28 && fromHeight == 24 {
		return true
	}

	return toHeight-fromHeight < 2
}

func neighbors(pos Pos) []Pos {
	return []Pos{
		{pos.X + 1, pos.Y + 1},
		{pos.X + 1, pos.Y - 1},
		{pos.X - 1, pos.Y + 1},
		{pos.X - 1, pos.Y - 1},
		{pos.X + 1, pos.Y},
		{pos.X - 1, pos.Y},
		{pos.X, pos.Y + 1},
		{pos.X, pos.Y - 1},
	}
}

func filterNeighbors(neighbors []Pos, visited [][]int) []Pos {
	filtered := []Pos{}
	for _, neighbor := range neighbors {
		if neighbor.X < 0 || len(visited[0])-1 < neighbor.X {
			continue
		}
		if neighbor.Y < 0 || len(visited)-1 < neighbor.Y {
			continue
		}
		if visited[neighbor.Y][neighbor.X] == 1 {
			continue
		}
		filtered = append(filtered, neighbor)
	}
	return filtered
}

func subset(terrain [][]int, domain []Pos) []int {
	subset := make([]int, len(domain))
	for i, pos := range domain {
		subset[i] = terrain[pos.Y][pos.X]
	}

	return subset
}

func path(domains [][]Pos) []Pos {
	path := []Pos{}

	return path
}

func printTerrain(terrain, visited [][]int) {
	fmt.Print("\033[H\033[2J")
	for y, row := range terrain {
		for x, col := range row {
			switch col {
			case -14:
				fmt.Print(color.InBold(color.InGreenOverWhite("S")))
			case -28:
				fmt.Print(color.InBold(color.InRedOverWhite("E")))
			default:
				if visited[y][x] == 1 {
					fmt.Print(color.InCyan(string(char(col))))
				} else {
					fmt.Print(color.InWhite(string(char(col))))
				}
			}

		}
		fmt.Println()
	}
}
