package medium

import (
	"testing"
)

func TestReverseInteger1(t *testing.T) {

	// Arange
	number := 123

	// Act
	result := reverse(number)
	should := 321

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestReverseInteger2(t *testing.T) {

	// Arange
	number := -123

	// Act
	result := reverse(number)
	should := -321

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestReverseInteger3(t *testing.T) {

	// Arange
	number := 120

	// Act
	result := reverse(number)
	should := 21

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestReverseInteger4(t *testing.T) {

	// Arange
	number := 1534236469

	// Act
	result := reverse(number)
	should := 0

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestReverseInteger5(t *testing.T) {

	// Arange
	number := -2147483648

	// Act
	result := reverse(number)
	should := 0

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestReverseInteger6(t *testing.T) {

	// Arange
	number := 1563847412

	// Act
	result := reverse(number)
	should := 0

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
