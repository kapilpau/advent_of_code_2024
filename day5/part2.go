package day5

import (
	"bufio"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

func part2() int {

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
			if !checkRules(rules, pages) {
				correct_updates = append(correct_updates, pages)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, pages := range correct_updates {
		count += reorderPages(pages, rules)[(len(pages)-1)/2]
	}

	return count
}

func reorderPages(pages []int, rules map[int][]int) []int {
	new_pages := make([]int, len(pages))

	copy(new_pages, pages)
	slices.SortStableFunc(new_pages, func(i, j int) int {
		if rule, ok := rules[i]; ok && slices.Contains(rule, j) {
			return 0
		}
		return -1
	})

	return new_pages
}
