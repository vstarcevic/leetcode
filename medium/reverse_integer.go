// Reverse Integer
// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:
// Input: x = 123
// Output: 321

// Example 2:
// Input: x = -123
// Output: -321
package medium

import "math"

func reverse(x int) int {

	result := 0
	isNegative := x < 0

	if isNegative {
		x *= -1
	}

	var length = int(math.Log10(float64(x)) + 1)

	if length > 10 {
		return 0
	}
	if length == 10 && x%10 > 2 {
		return 0
	}

	for i := 0; i < length; i++ {
		result = result*10 + x%10
		x /= 10
	}

	if result >= int(math.Pow(2, 31)) {
		return 0
	}

	if isNegative {
		result *= -1
	}

	return result

}
