package medium

import (
	"reflect"
	"testing"
)

func TestCombinationSum1(t *testing.T) {

	// Arange
	num := []int{2, 3, 6, 7}
	target := 7

	// Act
	result := combinationSum(num, target)

	should := [][]int{{2, 2, 3}, {7}}

	// Assert
	if !reflect.DeepEqual(result, should) {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result[0], should[0])
	}
}
