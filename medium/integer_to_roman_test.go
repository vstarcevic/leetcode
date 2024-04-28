package medium

import (
	"testing"
)

func TestIntToRoman1(t *testing.T) {

	// Arange
	num := 1994

	// Act
	result := intToRoman(num)
	should := "MCMXCIV"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestIntToRoman2(t *testing.T) {

	// Arange
	num := 58

	// Act
	result := intToRoman(num)
	should := "LVIII"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestIntToRoman3(t *testing.T) {

	// Arange
	num := 3

	// Act
	result := intToRoman(num)
	should := "III"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestIntToRoman4(t *testing.T) {

	// Arange
	num := 400

	// Act
	result := intToRoman(num)
	should := "CD"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
