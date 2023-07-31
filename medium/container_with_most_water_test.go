package medium

import (
	"testing"
)

func TestContainerWithMostWater1(t *testing.T) {

	// Arange
	containers := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	// Act
	result := maxArea(containers)
	should := 49

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
