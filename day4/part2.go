package day4

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

func part2() int {

	var grid [][]string

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path.Join(pwd, "./day4/input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	line_number := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		if grid == nil {
			grid = make([][]string, len(line))
		}
		grid[line_number] = line
		line_number++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	as := make([][]int, 0)

	for i, line := range grid {
		for j, char := range line {
			if char == "A" && i >= 1 && i < len(grid)-1 && j >= 1 && j < len(grid[0])-1 {
				as = append(as, []int{i, j})
			}
		}
	}

	for _, a := range as {
		if check_mas(a[0], a[1], grid) {
			count++
		}
	}

	return count
}

func check_mas(row, col int, grid [][]string) bool {
	count := 0

	if grid[row-1][col-1] == "M" && grid[row+1][col+1] == "S" {
		count++
	}

	if grid[row+1][col+1] == "M" && grid[row-1][col-1] == "S" {
		count++
	}

	if grid[row+1][col-1] == "M" && grid[row-1][col+1] == "S" {
		count++
	}

	if grid[row-1][col+1] == "M" && grid[row+1][col-1] == "S" {
		count++
	}

	return count == 2
}
