package easy

import (
	"testing"
)

func TestSqrt1(t *testing.T) {

	var x = 8

	should := 2
	result := mySqrt(x)

	if result != should {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}

}

func TestSqrt2(t *testing.T) {

	var x = 9

	should := 3
	result := mySqrt(x)

	if result != should {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}

}

func TestSqrt3(t *testing.T) {

	var x = 90

	should := 9
	result := mySqrt(x)

	if result != should {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}

}
