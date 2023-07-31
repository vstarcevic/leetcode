package medium

import (
	"testing"
)

func TestStringToInt1(t *testing.T) {

	// Arange
	word := "-123"

	// Act
	result := myAtoi(word)
	should := -123

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt2(t *testing.T) {

	// Arange
	word := "123"

	// Act
	result := myAtoi(word)
	should := 123

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt3(t *testing.T) {

	// Arange
	word := "words and 987"

	// Act
	result := myAtoi(word)
	should := 0

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt4(t *testing.T) {

	// Arange
	word := "987 and words"

	// Act
	result := myAtoi(word)
	should := 987

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt5(t *testing.T) {

	// Arange
	word := "-91283472332"

	// Act
	result := myAtoi(word)
	should := -2147483648

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt6(t *testing.T) {

	// Arange
	word := "   -42"

	// Act
	result := myAtoi(word)
	should := -42

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt7(t *testing.T) {

	// Arange
	word := "+-12"

	// Act
	result := myAtoi(word)
	should := 0

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt8(t *testing.T) {

	// Arange
	word := "9223372036854775808"

	// Act
	result := myAtoi(word)
	should := 2147483647

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt9(t *testing.T) {

	// Arange
	word := "   +0 123"

	// Act
	result := myAtoi(word)
	should := 0

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt10(t *testing.T) {

	// Arange
	word := "-5-"

	// Act
	result := myAtoi(word)
	should := -5

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt11(t *testing.T) {

	// Arange
	word := "00000-42a1234"

	// Act
	result := myAtoi(word)
	should := 0

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt12(t *testing.T) {

	// Arange
	word := "3.14159"

	// Act
	result := myAtoi(word)
	should := 3

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt13(t *testing.T) {

	// Arange
	word := "  -0012a42"

	// Act
	result := myAtoi(word)
	should := -12

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt14(t *testing.T) {

	// Arange
	word := "-13+8"

	// Act
	result := myAtoi(word)
	should := -13

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestStringToInt15(t *testing.T) {

	// Arange
	word := "18446744073709551617"

	// Act
	result := myAtoi(word)
	should := 2147483647

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
