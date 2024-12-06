package day6

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strings"
)

type Guard struct {
	icon           string
	row            int
	column         int
	originalIcon   string
	originalRow    int
	originalColumn int
}

func newGuard(icon string, row int, column int) Guard {
	return Guard{
		icon:           icon,
		row:            row,
		column:         column,
		originalIcon:   icon,
		originalRow:    row,
		originalColumn: column,
	}
}

func (g *Guard) reset() {
	g.icon = g.originalIcon
	g.row = g.originalRow
	g.column = g.originalColumn
}

func (g *Guard) move(grid [][]string, visitOnce bool) error {
	var nextRow, nextColumn int
	switch g.icon {
	case "^":
		nextRow = g.row - 1
		nextColumn = g.column
	case ">":
		nextColumn = g.column + 1
		nextRow = g.row
	case "v":
		nextRow = g.row + 1
		nextColumn = g.column
	case "<":
		nextColumn = g.column - 1
		nextRow = g.row
	}

	grid[g.row][g.column] = "X"
	if (nextRow < 0) || (nextColumn < 0) || (nextRow >= len(grid)) || (nextColumn >= len(grid[0])) {
		return nil
	} else if visitOnce && grid[nextRow][nextColumn] == "X" {
		return errors.New("already visited")
	} else if grid[nextRow][nextColumn] == "#" {
		g.turn()
		return g.move(grid, visitOnce)
	} else {
		g.row = nextRow
		g.column = nextColumn
		return g.move(grid, visitOnce)
	}
	return nil
}

func (g *Guard) turn() {
	switch g.icon {
	case "^":
		g.icon = ">"
	case ">":
		g.icon = "v"
	case "v":
		g.icon = "<"
	case "<":
		g.icon = "^"
	}
}

func part1() int {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path.Join(pwd, "./day6/input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	var guard Guard
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "^") {
			guard = newGuard("^", len(grid), strings.Index(line, "^"))
		}
		grid = append(grid, strings.Split(line, ""))
	}

	if guard == (Guard{}) {
		log.Fatal("no guards found")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if guard.move(grid, false) != nil {
		log.Fatal("cross")
	}
	return countXs(grid)
}

func countXs(grid [][]string) int {
	count := 0

	for _, row := range grid {
		for _, column := range row {
			if column == "X" {
				count++
			}
		}
	}

	return count
}
