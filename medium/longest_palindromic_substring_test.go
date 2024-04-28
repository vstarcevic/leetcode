package medium

import (
	"testing"
)

func TestLongestPalindromicSubstring1(t *testing.T) {

	// Arange
	str := "cbbd"
	// Act
	result := longestPalindrome(str)
	should := "bb"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestLongestPalindromicSubstring2(t *testing.T) {

	// Arange
	str := "babad"
	// Act
	result := longestPalindrome(str)
	should1 := "bab"
	should2 := "aba"

	// Assert
	if result != should1 && result != should2 {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v or %v`, result, should1, should2)
	}
}

func TestLongestPalindromicSubstring3(t *testing.T) {

	// Arange
	str := "aba"
	// Act
	result := longestPalindrome(str)
	should1 := "aba"

	// Assert
	if result != should1 {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should1)
	}
}
