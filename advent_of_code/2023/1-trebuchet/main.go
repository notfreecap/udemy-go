package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var digits = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func main() {
	data, err := os.ReadFile("data")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}

	result := 0
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}
		fd := findFirstDigit(line)
		ld := findLastDigit(line)

		log.Printf("Line %v - Fisrt digit: %v - Last digit: %v", i+1, fd, ld)
		result += int(fd*10 + ld)
	}

	fmt.Printf("The sum of all calibration values is: %v", result)
}

func findFirstDigit(line string) int {
	d := -1
	for i, c := range line {
		cIsNum, val := characterIsNumber(c)
		if cIsNum {
			d = val
			break
		}

		substr := line[i:]
		sIsNum, val := strIsNumber(substr, true)
		if sIsNum {
			d = val
			break
		}

	}
	return d
}

func findLastDigit(line string) int {
	d := -1
	for i := range line {
		lineIndex := len(line) - 1 - i
		c := rune(line[lineIndex])
		cIsNum, val := characterIsNumber(c)
		if cIsNum {
			d = val
			break
		}

		substr := line[:len(line)-i]
		sIsNum, val := strIsNumber(substr, false)
		if sIsNum {
			d = val
			break
		}
	}

	return d
}

func characterIsNumber(c rune) (isNumber bool, val int) {
	isNumber = unicode.IsNumber(c)
	if isNumber == true {
		val = int(c - '0')
	} else {
		val = -1
	}
	return
}

func strIsNumber(str string, searchPrefix bool) (isNumber bool, val int) {
	for key, value := range digits {
		if searchPrefix {
			if strings.HasPrefix(str, key) {
				isNumber = true
				val = value
				break
			}
		} else {
			if strings.HasSuffix(str, key) {
				isNumber = true
				val = value
				break
			}
		}
	}
	return
}
