package day2

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
	file, err := os.Open(path.Join(pwd, "./day2/input"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safe_count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		string_levels := strings.Split(scanner.Text(), " ")
		if is_level_safe(string_levels, -1) {
			safe_count++
		} else {
			for i := range string_levels {
				if is_level_safe(string_levels, i) {
					safe_count++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return safe_count
}
