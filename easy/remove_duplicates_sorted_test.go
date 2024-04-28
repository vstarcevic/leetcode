package easy

import (
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {

	// Arange
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} // Input array
	expectedNums := []int{0, 1, 2, 3, 4}        // Result can have elements after last one, does not matter {0, 1, 2, 3, 4, _, _, ...}

	// Act
	k := removeDuplicates(nums)

	// Assert
	if k != len(expectedNums) {
		t.Fatalf(`Not correct number of elements. Got %v, and it should be %v`, k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			t.Fatalf(`Result not as expected. Got %v, and it should be %v`, nums, expectedNums)
		}
	}
}

func TestRemoveDuplicates2(t *testing.T) {

	// Arange
	nums := []int{1, 2}         // Input array
	expectedNums := []int{1, 2} // Result can have elements after last one, does not matter {0, 1, 2, 3, 4, _, _, ...}

	// Act
	k := removeDuplicates(nums)

	// Assert
	if k != len(expectedNums) {
		t.Fatalf(`Not correct number of elements. Got %v, and it should be %v`, k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			t.Fatalf(`Result not as expected. Got %v, and it should be %v`, nums, expectedNums)
		}
	}
}
