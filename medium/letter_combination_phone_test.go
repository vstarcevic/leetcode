package medium

import (
	"testing"
)

func TestLetterCombinationPhone1(t *testing.T) {

	// Arange
	digits := "23"

	// Act
	result := letterCombinations(digits)
	should := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	// Assert
	for i := 0; i < len(result); i++ {
		if result[i] != should[i] {
			t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
		}
	}
}

func TestLetterCombinationPhone2(t *testing.T) {

	// Arange
	digits := ""

	// Act
	result := letterCombinations(digits)
	should := []string{}

	// Assert
	for i := 0; i < len(result); i++ {
		if result[i] != should[i] {
			t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
		}
	}

}

func TestLetterCombinationPhone3(t *testing.T) {

	// Arange
	digits := "2"

	// Act
	result := letterCombinations(digits)
	should := []string{"a", "b", "c"}

	// Assert
	for i := 0; i < len(result); i++ {
		if result[i] != should[i] {
			t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
		}
	}

}
