package day2

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2() int {
	filePath := "day2/reports_input.csv"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		var row []int
		for _, v := range values {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Failed to convert '%s' to int: %v", v, err)
			}
			row = append(row, num)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	safeReports := 0

	for _, record := range data {
		safe := isSafe(record)
		if safe {
			safeReports += 1
		}
	}

	return safeReports
}

func isSafe(record []int) bool {
	direction := 0

	for i, v := range record {
		if i == 1 {
			if v > record[0] {
				direction = 1
			}
			if v < record[0] {
				direction = -1
			}
			if v == record[0] {
				return false
			}
			if math.Abs(float64(v-record[i-1])) > 3 || math.Abs(float64(v-record[i-1])) < 1 {
				return false
			}
		}

		if i > 1 {
			if direction == 0 {
				return false
			}
			if math.Abs(float64(v-record[i-1])) > 3 || math.Abs(float64(v-record[i-1])) < 1 {
				return false
			}
			if direction == -1 && v > record[i-1] {
				return false
			}
			if direction == 1 && v < record[i-1] {
				return false
			}
		}
	}

	return true
}
