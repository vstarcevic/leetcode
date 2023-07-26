package easy

import (
	"testing"
)

func TestValidParentheses1(t *testing.T) {

	// Arange
	word := "{}"

	// Act
	result := isValid(word)
	should := true

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestValidParentheses2(t *testing.T) {

	// Arange
	word := "{]"

	// Act
	result := isValid(word)
	should := false

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestValidParentheses3(t *testing.T) {

	// Arange
	word := "]"

	// Act
	result := isValid(word)
	should := false

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}

}

func TestValidParentheses4(t *testing.T) {

	// Arange
	word := "(])"

	// Act
	result := isValid(word)
	should := false

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}

}
