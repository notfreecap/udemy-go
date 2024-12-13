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
	data, err := os.ReadFile("data")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	sum := 0
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := re.FindAllStringSubmatch(line, -1)
		iterateMatches(&sum, matches)
	}

	fmt.Println("total", sum)
}

func iterateMatches(sum *int, matches [][]string) {
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		*sum += a * b
	}
}
