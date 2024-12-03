package day3

import (
	"bufio"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
)

func part2() int {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path.Join(pwd, "./day3/input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	enabled := true
	number_extractor := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, match := range number_extractor.FindAllString(line, -1) {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else if parts := number_extractor.FindStringSubmatch(match); enabled && len(parts) > 2 {
				num1, _ := strconv.Atoi(parts[1])
				num2, _ := strconv.Atoi(parts[2])
				sum += num1 * num2
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
