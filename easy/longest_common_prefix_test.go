package easy

import (
	"testing"
)

func TestLongestCommonPrefix1(t *testing.T) {

	// Arange
	var arr = []string{"flower", "flow", "flight"}

	// Act
	result := longestCommonPrefix(arr)
	should := "fl"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestLongestCommonPrefix2(t *testing.T) {

	// Arange
	var arr = []string{"ab", "a"}

	// Act
	result := longestCommonPrefix(arr)
	should := "a"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestLongestCommonPrefix3(t *testing.T) {

	// Arange
	var arr = []string{"flower", "flower", "flower", "flower"}

	// Act
	result := longestCommonPrefix(arr)
	should := "flower"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestLongestCommonPrefix4(t *testing.T) {

	// Arange
	var arr = []string{"cir", "car"}

	// Act
	result := longestCommonPrefix(arr)
	should := "c"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
