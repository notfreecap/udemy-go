package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("data2")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}

	sum := 0
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}
		isFirstLine := i == 0
		checkValidity(line, isFirstLine, &sum)
		// iterateMatches(&sum, line)
	}

	fmt.Println("total", sum)
}

func iterateMatches(sum *int, line string) {
	fmt.Println(line)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		*sum += a * b
	}
}

func checkValidity(line string, isFirstLine bool, sum *int) {
	re := regexp.MustCompile(`[^d]*(?:do\(\)|'?don't\(\))`)
	matches := re.FindAllString(line, -1)
	for i, match := range matches {
		if i == 0 && isFirstLine {
			recheck := regexp.MustCompile(`do\(\)|don't\(\)`)
			result := recheck.Split(match, -1)
			if len(result) > 0 {
				iterateMatches(sum, result[0])
			}
			continue
		}
		if strings.Contains(match, "don't()") {
			continue
		}
		iterateMatches(sum, match)
	}
}
