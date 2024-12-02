package day2

import (
	"bufio"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

func is_level_safe(string_levels []string, level_to_skip int) bool {
	reduction := 0
	if level_to_skip > -1 {
		reduction = 1
	}
	levels := make([]int, len(string_levels)-reduction)

	for i, level := range string_levels {
		if i == level_to_skip {
			continue
		}

		level_pointer := i
		if level_to_skip > -1 && i > level_to_skip {
			level_pointer--
		}
		int_level, _ := strconv.Atoi(level)

		levels[level_pointer] = int_level

		if level_pointer > 0 {
			diff := abs(levels[level_pointer] - levels[level_pointer-1])
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}

	direction := 1
	if levels[0] > levels[1] {
		direction = -1
	}

	for i := 0; i < len(levels)-1; i++ {
		curr := levels[i]
		next := levels[i+1]

		if (direction == 1 && curr > next) || (direction == -1 && curr < next) {
			return false
		}
	}
	return true
}

func part1() int {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path.Join(pwd, "./day2/input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safe_count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if is_level_safe(strings.Split(scanner.Text(), " "), -1) {
			safe_count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return safe_count
}

func abs(num int) int {
	return int(math.Abs(float64((num))))
}
