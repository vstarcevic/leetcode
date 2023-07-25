package easy

import (
	"leetcode/pkg"
	"testing"
)

func TestTwoSum1(t *testing.T) {

	var i = []int{2, 7, 11, 15}
	var n = 9

	var should []int = []int{0, 1}
	result := twoSum(i, n)

	if !pkg.Equal(should, result) {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}

}

func TestTwoSum2(t *testing.T) {

	var i = []int{3, 3}
	var n = 6

	var should []int = []int{0, 1}
	result := twoSum(i, n)

	if !pkg.Equal(should, result) {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}
}

func TestTwoSum3(t *testing.T) {

	var i = []int{3, 2, 4}
	var n = 6

	var should []int = []int{1, 2}
	result := twoSum(i, n)

	if !pkg.Equal(should, result) {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}
}
