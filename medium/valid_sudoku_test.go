package medium

import (
	"testing"
)

func TestValidSudoku1(t *testing.T) {

	// Arange
	arr := [][]byte{
		{53, 51, 46, 46, 55, 46, 46, 46, 46},
		{54, 46, 46, 49, 57, 53, 46, 46, 46},
		{46, 57, 56, 46, 46, 46, 46, 54, 46},
		{56, 46, 46, 46, 54, 46, 46, 46, 51},
		{52, 46, 46, 56, 46, 51, 46, 46, 49},
		{55, 46, 46, 46, 50, 46, 46, 46, 54},
		{46, 54, 46, 46, 46, 46, 50, 56, 46},
		{46, 46, 46, 52, 49, 57, 46, 46, 53},
		{46, 46, 46, 46, 56, 46, 46, 55, 57}}

	// Act
	result := isValidSudoku(arr)
	should := true

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
