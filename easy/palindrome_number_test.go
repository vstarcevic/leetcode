package easy

import (
	"testing"
)

func TestPalindromeNumber1(t *testing.T) {

	// Arange
	var i = 11

	// Act
	result := isPalindrome(i)

	// Assert
	if result == false {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, true)
	}
}

func TestPalindromeNumber2(t *testing.T) {

	// Arange
	var i = 102

	// Act
	result := isPalindrome(i)

	// Assert
	if result == true {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, false)
	}
}

func TestPalindromeNumber3(t *testing.T) {

	// Arange
	var i = -121

	// Act
	result := isPalindrome(i)

	// Assert
	if result == true {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, false)
	}
}

func TestPalindromeNumber4(t *testing.T) {

	// Arange
	var i = -0

	// Act
	result := isPalindrome(i)

	// Assert
	if result == false {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, false)
	}
}
