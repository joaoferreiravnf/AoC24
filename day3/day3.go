package day3

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day3() int {
	filePath := "day3/corrupted_input"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	defer file.Close()

	totalSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		mulPattern := `mul\(\d{1,3},\d{1,3}\)`

		rePattern := regexp.MustCompile(mulPattern)

		patternMatches := rePattern.FindAllString(line, -1)

		for _, match := range patternMatches {
			patterNumber := `\d{1,3}`

			reNumber := regexp.MustCompile(patterNumber)
			numberMatches := reNumber.FindAllString(match, -1)

			numA, _ := strconv.Atoi(numberMatches[0])
			numB, _ := strconv.Atoi(numberMatches[1])

			numAB := numA * numB
			totalSum += numAB
		}
	}
	return totalSum
}
