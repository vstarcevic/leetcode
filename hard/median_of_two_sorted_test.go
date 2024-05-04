package hard

import (
	"testing"
)

func Test_Median_Of_Two_Sorted(t *testing.T) {

	// Arange
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5}
	// Act
	result := findMedianSortedArrays(nums1, nums2)

	// Assert
	if result-3 > 0.1 {
		t.Fatalf(`Result not as expected. Got %v`, result)
	}

}
