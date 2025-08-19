package main

import "testing"

func TestValidISBN(t *testing.T) {
	validISBN := "0-306-40615-2"

	want := true

	result := TestISBN(validISBN)

	if result != want {
		t.Errorf(`TestISBN("0-306-40615-2") = %t, want match for %t, nil`, result, want)
	}
}

func TestFailISBN(t *testing.T) {
	validISBN := "1234567890"

	want := false

	result := TestISBN(validISBN)

	if result != want {
		t.Errorf(`TestISBN("1234567890") = %t, want match for %t, nil`, result, want)
	}
}
