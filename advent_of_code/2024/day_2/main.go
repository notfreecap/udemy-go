package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("data")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}

	safeLines := 0
	safeLinesNew := 0
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		// if evaluateSafety(line, false) {
		// 	safeLines++
		// }

		if evaluateSafety(line, true) {
			safeLinesNew++
		}
	}
	log.Printf("Safe reports: %v", safeLines)
	log.Printf("Safe reports v2: %v", safeLinesNew)

}

func evaluateSafety(line string, tolerateFail bool) bool {
	numbers := getNumbers(line)
	log.Print("line->", line)
	if numbers[0]-numbers[1] > 0 {
		return validateDiff(numbers, true, tolerateFail)
	} else {
		return validateDiff(numbers, false, tolerateFail)
	}
}

func validateDiff(numbers []int, isPositive bool, tolerateFail bool) bool {
	log.Print("num->", numbers)
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i] - numbers[i+1]
		if (isPositive && (diff <= 0 || diff > 3)) || (!isPositive && (diff >= 0 || diff < -3)) {
			if tolerateFail {
				return validateDiff(append(numbers[:i], numbers[i+1:]...), isPositive, false)
			}
			return false
		}
	}
	return true
}

func getNumbers(line string) []int {
	var numbers []int
	str := strings.Split(line, " ")
	for _, strNum := range str {
		val, _ := strconv.Atoi(strNum)
		numbers = append(numbers, val)
	}

	return numbers
}
