package medium

import (
	"testing"
)

func TestSearchingChallenge(t *testing.T) {

	// Arange
	words := "Today is the greatest day ever!"

	// Act
	result := SearchingChallenge(words)
	should := "greatest"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
