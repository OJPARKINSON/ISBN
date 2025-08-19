package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	startingISBN := "0-306-40615-2"

	result := TestISBN(startingISBN)

	if result {
		fmt.Println("success")
	} else {
		fmt.Println("error")
	}
}

func TestISBN(startingISBN string) bool {
	match, err := regexp.MatchString(`\d-?\d{1,7}-?\d{1,5}-?(\d|X)`, startingISBN)
	if err != nil || !match {
		log.Fatal("Failed to parse")
	}

	split := strings.Split(startingISBN, "")

	count := 0
	countIndex := 0
	for i := 0; i < len(split); i++ {
		if i == len(split)-1 {
			continue
		} else if split[i] == "-" {
			continue
		} else {
			checkNum, _ := strconv.Atoi(split[i])

			count += checkNum * (countIndex + 1)
			countIndex++
		}
	}

	checkSum := count % 11

	lastDigit := 0
	lastIndex := split[len(split)-1]
	if lastIndex == "X" {
		lastDigit = 10
	} else {
		lastDigit, _ = strconv.Atoi(lastIndex)

	}

	return checkSum == lastDigit
}
