package main

import (
	"fmt"
	"os"
)

type Pos struct {
	X int
	Y int
}

func eq(a, b Pos) bool {
	return (a.X == b.X) && (a.Y == b.Y)
}

func main() {
	raw, err := os.ReadFile("./12/puzzle.input")
	if err != nil {
		panic(err.Error())
	}

	terrain, start, goal := parseTerrain(raw)
	shortest := search(terrain, []Pos{start}, goal)
	fmt.Println(shortest)

	shortest2 := search(terrain, getZeroes(terrain), goal)
	fmt.Println(shortest2)

}

func parseTerrain(raw []byte) (terrain [][]int, start, goal Pos) {
	terrain = append(terrain, []int{})

	x, y := 0, 0
	for _, c := range raw[:len(raw)-1] {
		if c == '\n' {
			y++
			x = 0
			terrain = append(terrain, []int{})
		} else {
			if c == 'S' {
				start = Pos{x, y}
				terrain[y] = append(terrain[y], 0)
			} else if c == 'E' {
				goal = Pos{x, y}
				terrain[y] = append(terrain[y], 25)
			} else {
				terrain[y] = append(terrain[y], height(c))
			}
			x++
		}
	}

	return terrain, start, goal
}

func height(c byte) int {
	return int(c) - 97
}

func char(i int) byte {
	return byte(i + 97)
}

func search(terrain [][]int, domain []Pos, goal Pos) int {
	shortest := len(terrain) * len(terrain[0])

	for _, pos := range domain {
		current := len(terrain) * len(terrain[0])
		distances := map[Pos]int{pos: 0}
		toVisit := []Pos{pos}

		for 0 < len(toVisit) {
			this := toVisit[0]
			toVisit = toVisit[1:]

			if eq(this, goal) && distances[this] < current {
				current = distances[this]
			}

			for _, neighbor := range neighbors(this, terrain) {
				_, visited := distances[neighbor]

				if !visited {
					toVisit = append(toVisit, neighbor)
					distances[neighbor] = distances[this] + 1
				}
			}
		}

		if current < shortest {
			shortest = current
		}
	}

	return shortest
}

func neighbors(pos Pos, terrain [][]int) []Pos {
	neighbors := make([]Pos, 0, 4)

	for _, neighbor := range []Pos{
		{pos.X + 1, pos.Y},
		{pos.X - 1, pos.Y},
		{pos.X, pos.Y + 1},
		{pos.X, pos.Y - 1},
	} {
		if neighbor.X < 0 || len(terrain[0])-1 < neighbor.X {
			continue
		}
		if neighbor.Y < 0 || len(terrain)-1 < neighbor.Y {
			continue
		}
		if terrain[neighbor.Y][neighbor.X]-terrain[pos.Y][pos.X] <= 1 {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func getZeroes(terrain [][]int) []Pos {
	zeroes := []Pos{}
	for y, row := range terrain {
		for x, value := range row {
			if value == 0 {
				zeroes = append(zeroes, Pos{x, y})
			}
		}
	}

	return zeroes
}
