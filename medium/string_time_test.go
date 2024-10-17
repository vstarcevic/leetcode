// Given one sentence, like "Today is the greatest day ever!"
// Find the word that have the most occuring letters.
// E.g greatest have two e's two t's and it's before ever which have only two e's
package medium

import (
	"testing"
)

func TestStringTime(t *testing.T) {

	// Arange
	words := []string{"10:00am", "11:45pm", "5:00am", "12:01am"}

	// Act
	result := StringTimeMinDiff(words)
	should := "16"

	// Assert
	if result != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should)
	}
}
