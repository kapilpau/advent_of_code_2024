package day5

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

func part1() int {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path.Join(pwd, "./day5/input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	correct_updates := make([][]int, 0)
	rules := make(map[int][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			first, _ := strconv.Atoi(strings.Split(line, "|")[0])
			second, _ := strconv.Atoi(strings.Split(line, "|")[1])
			if _, ok := rules[first]; !ok {
				rules[first] = make([]int, 0)
			}
			rules[first] = append(rules[first], second)
		} else if strings.Contains(line, ",") {
			pages := toIntSlice(line)
			if checkRules(rules, pages) {
				correct_updates = append(correct_updates, pages)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, pages := range correct_updates {
		count += pages[(len(pages)-1)/2]
	}

	return count
}

func toIntSlice(s string) []int {
	parts := strings.Split(s, ",")
	ints := make([]int, len(parts))
	for i, part := range parts {
		ints[i], _ = strconv.Atoi(part)
	}
	return ints
}

func checkRules(rules map[int][]int, pages []int) bool {
	for first, seconds := range rules {
		for _, second := range seconds {
			first, second := arrayPositions(pages, first, second)
			if first != -1 && second != -1 && second < first {
				return false
			}
		}
	}
	return true
}

func arrayPositions(arr []int, val1, val2 int) (int, int) {
	first := -1
	second := -1
	for i, v := range arr {
		if v == val1 {
			first = i
		}
		if v == val2 {
			second = i
		}
		if first != -1 && second != -1 {
			return first, second
		}
	}
	return first, second
}
