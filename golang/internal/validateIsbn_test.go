package internal

import (
	"testing"
)

func TestPassISBN(t *testing.T) {
	var passISBNS = []string{
		"0-306-40615-2", // classic valid ISBN
		"0-321-14653-0", // another real valid ISBN
		"0-9752298-0-X", // checksum X case
		"0306406152",    // same as first, without hyphens
		"123456789X",    // edge case: ends with X
		"1111111111",
	}

	for _, isbn := range passISBNS {
		t.Run(isbn, func(t *testing.T) {
			ans := ValidateISBN(isbn)
			if !ans {
				t.Errorf("got %t, wanted %t, for %s", ans, true, isbn)
			}
		})
	}
}

func TestFailISBN(t *testing.T) {
	var failISBNS = []string{
		"1234567890",    // checksum doesn't match
		"0-306-40615-9", // wrong checksum digit
		"ABCDEFGHIJ",    // non-numeric characters
		"123456789",     // too short (only 9 digits)
		"123456789XX",   // too long (11 characters)
		"",              // empty string
		"0-9752298-0-Y", // invalid last char (Y is not allowed)
	}

	for _, isbn := range failISBNS {
		t.Run(isbn, func(t *testing.T) {
			ans := ValidateISBN(isbn)
			if ans {
				t.Errorf("got %t, wanted %t, for %s", ans, false, isbn)
			}
		})
	}
}
