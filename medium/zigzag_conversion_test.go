package medium

import (
	"testing"
)

func TestZigZagConversion1(t *testing.T) {

	// Arange
	word := "PAYPALISHIRING"
	rows := 3

	// Act
	result := convert(word, rows)
	should := "PAHNAPLSIIGYIR"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestZigZagConversion2(t *testing.T) {

	// Arange
	word := "PAYPALISHIRING"
	rows := 4

	// Act
	result := convert(word, rows)
	should := "PINALSIGYAHRPI"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}

func TestZigZagConversion3(t *testing.T) {

	// Arange
	word := "A"
	rows := 1

	// Act
	result := convert(word, rows)
	should := "A"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
