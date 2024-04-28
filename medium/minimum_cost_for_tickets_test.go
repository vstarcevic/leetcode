package medium

import (
	"testing"
)

func TestMinCostTickets1(t *testing.T) {

	// Arange
	arr := []int{1, 4, 6, 7, 8, 20}
	cost := []int{2, 7, 15}

	// Act
	result := mincostTickets(arr, cost)
	should := 11

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
