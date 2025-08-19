package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	startingISBN := "123456789XX"

	result := ValidateISBN(startingISBN)

	if result {
		fmt.Println("success")
	} else {
		fmt.Println("error")
	}
}

func ValidateISBN(startingISBN string) bool {
	match, err := regexp.MatchString(`((\d-)|\d{1,10})\d{1,10}?-\d{0,5}?-(\d{1}|X)|(\d{1,10}(X|\d{1}))$`, startingISBN)
	if err != nil || !match {
		fmt.Println(err)
		return false
	}

	splitISBN := strings.Split(startingISBN, "")

	count := 0
	countIndex := 0
	for i := 0; i < len(splitISBN); i++ {
		if i == len(splitISBN)-1 {
			continue
		} else if splitISBN[i] == "-" {
			continue
		} else {
			num, _ := strconv.Atoi(splitISBN[i])

			count += num * (countIndex + 1)
			countIndex++
		}

		fmt.Println("count: ", count)
	}

	numForCheck := count % 11
	fmt.Println(numForCheck)

	checkSum := 0
	lastIndex := splitISBN[len(splitISBN)-1]
	if lastIndex == "X" {
		checkSum = 10
	} else {
		checkSum, _ = strconv.Atoi(lastIndex)

	}
	fmt.Println(checkSum)

	return checkSum == numForCheck
}
