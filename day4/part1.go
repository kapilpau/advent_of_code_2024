package day4

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

const SEARCH_STRING = "XMAS"

func part1() int {

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

	xs := make([][]int, 0)

	for i, line := range grid {
		for j, char := range line {
			if char == "X" {
				xs = append(xs, []int{i, j})
			}
		}
	}

	for _, x := range xs {
		count += xmas_search_around(x[0], x[1], grid)
	}

	return count
}

func xmas_search_around(row, col int, grid [][]string) int {
	count := 0

	min_row := row - 1
	if min_row < 0 {
		min_row = 0
	}

	max_row := row + 1
	if max_row > len(grid)-1 {
		max_row = len(grid) - 1
	}

	min_col := col - 1
	if min_col < 0 {
		min_col = 0
	}

	max_col := col + 1
	if max_col > len(grid[0])-1 {
		max_col = len(grid[0]) - 1
	}

	for i := min_row; i <= max_row; i++ {
		for j := min_col; j <= max_col; j++ {
			if grid[i][j] == "M" {
				A_i := i + (i - row)
				A_j := j + (j - col)

				if A_i < 0 || A_i > len(grid)-1 || A_j < 0 || A_j > len(grid[0])-1 {
					continue
				}

				if grid[A_i][A_j] == "A" {
					S_i := A_i + (A_i - i)
					S_j := A_j + (A_j - j)

					if S_i < 0 || S_i > len(grid)-1 || S_j < 0 || S_j > len(grid[0])-1 {
						continue
					}

					if grid[S_i][S_j] == "S" {
						count++
					}
				}
			}
		}
	}

	return count
}
