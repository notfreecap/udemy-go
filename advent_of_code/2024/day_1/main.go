package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("data")
	if err != nil {
		log.Fatal("Unable to read file", err)
	}

	var numbersLeft []int
	var numbersRight []int
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "   ")
		numbersLeft = append(numbersLeft, getNumber(parts[0]))
		numbersRight = append(numbersRight, getNumber(parts[1]))
	}

	sort.Ints(numbersLeft)
	sort.Ints(numbersRight)
	calcDistance(numbersLeft, numbersRight)
	calcSimilarityScore(numbersLeft, numbersRight)
}

func getNumber(strNumb string) int {
	num, err := strconv.Atoi(strNumb)
	if err != nil {
		num = 0
	}
	return num
}

func calcDistance(numbersLeft []int, numbersRight []int) {
	var total int = 0
	for i, leftNumber := range numbersLeft {
		total += int(math.Abs(float64(leftNumber - numbersRight[i])))
		//log.Printf("left: %v - right: %v", leftNumber, numbersRight[i])
	}
	log.Printf("Total distance: %v", total)
}

func calcSimilarityScore(numbersLeft []int, numbersRight []int) {
	groupedLeft := groupNumbers(numbersLeft)
	groupedRight := groupNumbers(numbersRight)

	score := 0
	for key, val := range groupedLeft {
		//log.Printf("number %v - times %v", key, val)
		score += (key * val) * groupedRight[key]
	}
	log.Printf("Similarity score %v", score)
}

func groupNumbers(numbers []int) map[int]int {
	grouped := make(map[int]int)
	for _, num := range numbers {
		grouped[num]++
	}
	return grouped
}
