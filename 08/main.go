package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, err := os.ReadFile("./08/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	terrain := NewTerrain(in)

	fmt.Println(terrain)

	calculateVisibility(terrain)
	total := countVisibleTrees(terrain)

	fmt.Printf("Part 1: There are %d visible trees.", total)
}

type Exposure = int

const (
	Hidden  Exposure = -1
	Unknown Exposure = 0
	Visible Exposure = 1
)

type Tree struct {
	Height   int
	Exposure Exposure
}

type Terrain struct {
	trees [][]Tree
}

func NewTerrain(in []byte) *Terrain {
	rows := strings.Split(strings.TrimSpace(string(in)), "\n")

	trees := [][]Tree{}
	for y, row := range rows {
		trees = append(trees, []Tree{})
		for _, t := range row {
			height, _ := strconv.Atoi(string(t))
			trees[y] = append(trees[y], Tree{Height: height})
		}
	}

	return &Terrain{
		trees: trees,
	}
}

func (t *Terrain) String() string {
	s := "Trees\n"
	for _, row := range t.trees {
		for _, val := range row {
			s += fmt.Sprint(val.Height)
		}
		s += "\n"
	}

	s += "\nVisibility\n"
	for _, row := range t.trees {
		for _, val := range row {
			s += fmt.Sprint(val.Exposure)
		}
		s += "\n"
	}

	return s
}

func (t *Terrain) Width() int {
	return len(t.trees[0])
}

func (t *Terrain) Height() int {
	return len(t.trees)
}

func (t *Terrain) Tree(x, y int) *Tree {
	if x < 0 || t.Width() <= x || y < 0 || t.Height() <= y {
		return nil
	}

	return &t.trees[y][x]
}

func (t *Terrain) LeftOf(x, y int) []Tree {
	trees := []Tree{}
	for i := x - 1; i >= 0; i-- {
		trees = append(trees, *t.Tree(i, y))
	}
	return trees
}

func (t *Terrain) RightOf(x, y int) []Tree {
	trees := []Tree{}
	for i := x + 1; i < t.Width(); i++ {
		trees = append(trees, *t.Tree(i, y))
	}
	return trees
}

func (t *Terrain) Above(x, y int) []Tree {
	trees := []Tree{}
	for i := y - 1; i >= 0; i-- {
		trees = append(trees, *t.Tree(x, i))
	}
	return trees
}

func (t *Terrain) Below(x, y int) []Tree {
	trees := []Tree{}
	for i := y + 1; i < t.Height(); i++ {
		trees = append(trees, *t.Tree(x, i))
	}
	return trees
}

func calculateVisibility(t *Terrain) {
	// Mark all of the edge trees as visible
	for x := 0; x < t.Width(); x++ {
		for y := 0; y < t.Height(); y++ {
			if x == 0 || x == t.Width()-1 || y == 0 || y == t.Height()-1 {
				t.Tree(x, y).Exposure = Visible
			} else {
				height := t.Tree(x, y).Height
				if allShorter(t.LeftOf(x, y), height) || allShorter(t.RightOf(x, y), height) || allShorter(t.Above(x, y), height) || allShorter(t.Below(x, y), height) {
					t.Tree(x, y).Exposure = Visible
				}
			}
		}
	}

	fmt.Println(t)
}

func allShorter(trees []Tree, height int) bool {
	for _, tree := range trees {
		if height <= tree.Height {
			return false
		}
	}
	return true
}

func countVisibleTrees(t *Terrain) int {
	count := 0
	for x := 0; x < t.Width(); x++ {
		for y := 0; y < t.Height(); y++ {
			if t.Tree(x, y).Exposure == Visible {
				count++
			}
		}
	}

	return count
}
