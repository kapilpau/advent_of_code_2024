package day6

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

func part2() int {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path.Join(pwd, "./day6/exampleinput"))
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
			guard = Guard{
				icon:   "^",
				row:    len(grid),
				column: strings.Index(line, "^"),
			}
		}
		grid = append(grid, strings.Split(line, ""))
	}

	if guard == (Guard{}) {
		log.Fatal("no guards found")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	for i, row := range grid {
		for j, column := range row {
			if column != "#" {
				grid[i][j] = "#"
				if guard.move(grid, true) != nil {
					count++
					guard.reset()
				}
				grid[i][j] = "."
			}
		}
	}
	// guard.move(grid, false)
	return count
}
