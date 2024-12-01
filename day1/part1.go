package day1

import (
	"bufio"
	"log"
	"math"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func part1() int {
	left := make([]int, 0)
	right := make([]int, 0)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path.Join(pwd, "./day1/input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		if num, err := strconv.Atoi(parts[0]); err == nil {
			left = append(left, int(num))
		}

		if num, err := strconv.Atoi(parts[1]); err == nil {
			right = append(right, int(num))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(right, func(i, j int) bool {
		return right[i] > right[j]
	})

	sort.Slice(left, func(i, j int) bool {
		return left[i] > left[j]
	})

	distance := 0

	for i := range right {
		distance += abs(left[i] - right[i])
	}

	return distance
}

func abs(num int) int {
	return int(math.Abs(float64((num))))
}
