package internal

import (
	"regexp"
	"strconv"
	"strings"
)

func ValidateISBN(input string) bool {
	match, _ := regexp.MatchString(`^[0-9Xx\s-.:]+$`, input)
	if !match {
		return false
	}

	normalisedISBN := normaliseISBN(input)

	isbnFormat := detectFormat(normalisedISBN)

	if isbnFormat == "invalid" {
		return false
	} else if isbnFormat == "ISBN-10" {
		return ValidateISBN10(normalisedISBN)
	} else {
		return ValidateISBN13(normalisedISBN)
	}
}

func ValidateISBN10(isbn10 string) bool {
	sum := 0

	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(isbn10[i]))
		sum += digit * (i + 1)
	}

	checksum := sum % 11
	lastChar := string(isbn10[9])

	if lastChar == "X" {
		return checksum == 10
	} else {
		lastDigit, _ := strconv.Atoi(lastChar)
		return checksum == lastDigit

	}
}

func ValidateISBN13(isbn13 string) bool {
	sum := 0

	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(isbn13[i]))
		if i%2 == 0 {
			sum += digit * 1
		} else {
			sum += digit * 3
		}
	}

	checksum := (10 - (sum % 10)) % 10
	lastDigit, _ := strconv.Atoi(string(isbn13[12]))
	return checksum == lastDigit
}

func normaliseISBN(input string) string {
	re := regexp.MustCompile(`^ISBN-1[03]:\s*`)
	cleaned := re.ReplaceAllString(input, "")

	re = regexp.MustCompile(`[\s\-\.]`)
	cleaned = re.ReplaceAllString(cleaned, "")

	return strings.ToUpper(cleaned)
}

func detectFormat(normalised string) string {
	switch len(normalised) {
	case 10:
		return "ISBN-10"
	case 13:
		if strings.HasPrefix(normalised, "978") || strings.HasPrefix(normalised, "979") {
			return "ISBN-13"
		}
		return "invalid"
	default:
		return "invalid"
	}

}
