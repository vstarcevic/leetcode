// Palindrome Number
// Given an integer x, return true if x is a
// palindrome, and false otherwise.

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

package easy

import "math"

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 {
		return false
	}

	if x%10 == 0 {
		return false
	}

	var digits int = int(math.Log10(float64(x))) + 1

	reversed := 0
	original := x

	for i := 0; i < digits; i++ {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == original

}
