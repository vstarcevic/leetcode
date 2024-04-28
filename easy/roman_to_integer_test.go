package easy

import (
	"testing"
)

func TestRomanToInteger1(t *testing.T) {

	var r = "III"

	should := 3
	result := romanToInt(r)

	if should != result {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}

}

func TestRomanToInteger2(t *testing.T) {

	var r = "MCMXCIV"

	should := 1994
	result := romanToInt(r)

	if should != result {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}

}

func TestRomanToInteger3(t *testing.T) {

	var r = "LVIII"

	should := 58
	result := romanToInt(r)

	if should != result {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, result, should)
	}

}
