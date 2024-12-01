package day1

import (
	"bufio"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func part2() int {
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
		return right[i] < right[j]
	})

	similarity := 0

	for _, num := range left {
		similarity += (num * count_occurences(right, num))
	}

	return similarity
}

func count_occurences(nums []int, num int) int {
	count := 0
	for _, n := range nums {
		if n == num {
			count++
		} else {
			if n > num {
				break
			}
		}
	}
	return count
}
