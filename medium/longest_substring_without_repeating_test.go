package medium

import (
	"testing"
)

func TestLongestSubstring1(t *testing.T) {

	// Arange
	word := "abcabcbb"

	// Act
	result := longestSubstringWithoutRepeatingCharacters(word)
	should := 3

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestLongestSubstring2(t *testing.T) {

	// Arange
	word := "aab"

	// Act
	result := longestSubstringWithoutRepeatingCharacters(word)
	should := 2

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
