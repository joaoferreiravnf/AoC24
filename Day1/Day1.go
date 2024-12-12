package Day1

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func Day1() int {
	csvRootFile := "/Users/joaoferreira/git_joao/trainingGo/AoC2.csv"

	var list1 []int
	var list2 []int

	csvFile, err := os.Open(csvRootFile)
	if err != nil {
		log.Fatalf("error opening CSV file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading CSV file: %v", err)
	}

	if len(records) < 2 {
		log.Fatalf("csv file is empty")
	}

	for _, record := range records {
		lis1Item, _ := strconv.Atoi(record[0])
		lis2Item, _ := strconv.Atoi(record[1])
		list1 = append(list1, lis1Item)
		list2 = append(list2, lis2Item)
	}

	totalDif := 0

	sort.Ints(list1)
	sort.Ints(list2)

	for i, v := range list1 {
		simDif := v - list2[i]
		totalDif += int(math.Abs(float64(simDif)))
	}

	return totalDif
}
