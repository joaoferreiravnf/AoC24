package day4

import (
	"bufio"
	"log"
	"os"
)

func Day4() int {
	filePath := "day4/word_search_input"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	defer file.Close()

	var input [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var lineString []string
		line := scanner.Text()
		for _, v := range line {
			lineString = append(lineString, string(v))
		}
		input = append(input, lineString)
	}

	count := 0

	for i, v := range input {
		for j, y := range v {
			if y == "X" {
				horiz := checkHorizontal(i, j, input)
				vert := checkVertical(i, j, input)
				diag := checkDiagonal(i, j, input)
				count += horiz + vert + diag
			}
		}
	}

	return count
}

func checkHorizontal(i, j int, input [][]string) int {
	if (j < len(input[i])-3 && input[i][j+1] == "M" && input[i][j+2] == "A" && input[i][j+3] == "S") ||
		(j > 2 && input[i][j-1] == "M" && input[i][j-2] == "A" && input[i][j-3] == "S") {
		return 1
	}
	return 0
}

func checkVertical(i, j int, input [][]string) int {
	if (i < len(input) - 3 && input[i+1][j] == "M" && input[i+2][j] == "A" && input[i+3][j] == "S") ||
		(i > 2 && input[i-1][j] == "M" && input[i-2][j] == "A" && input[i-3][j] == "S") {
		return 1
	}
	return 0
}

func checkDiagonal(i, j int, input [][]string) int {
	if
	return 0
}
