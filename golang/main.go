package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validateISBN(splitStr []string) {
	dashes := 0
	for i := 0; i < len(splitStr); i++ {
		if splitStr[i] == "-" {
			if splitStr[i+1] == "-" {
				panic("Invalid ISBN: consecutive dashes")
			}
			dashes++
		}
	}

	if dashes > 2 {
		panic("Invalid ISBN: too many dashes")
	}
}

func countISBN(ISBN string) int {
	splitStr := strings.Split(ISBN, "")

	validateISBN(splitStr)

	count := 0
	descSeq := 10
	for i := 0; i < len(splitStr); i++ {
		if splitStr[i] == "-" {
			splitStr = append(splitStr[:i], splitStr[i+1:]...)
			i -= 1
			continue
		} else {
			num, err := strconv.Atoi(splitStr[i])
			if err != nil {
				panic("Failed to parse str")
			}
			count += num * descSeq
			descSeq--
		}
	}

	return count

}

func main() {
	ISBN := "0-596-52068"

	count := countISBN(ISBN)

	remainder := count % 11

	fmt.Println(11 - remainder)

}
