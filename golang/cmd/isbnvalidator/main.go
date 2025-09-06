package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/OJPARKINSON/ISBN/golang/internal"
)

func main() {
	argsWithoutProg := os.Args[1:]

	file, err := os.Open(argsWithoutProg[0])
	if err != nil {
		log.Panic("Failed to get file")
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Panic("Failed to read the csv file")
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		result := internal.ValidateISBN(record[0])

		expectedResult, err := strconv.ParseBool(record[1])
		if err != nil {
			log.Fatal("failed to parse expected value: ", err)
		}

		if result == expectedResult {
			fmt.Println(record[0], " - success")
		} else {
			fmt.Println(record[0], " - error: result: ", result, " expected: ", expectedResult)
		}
	}

}
