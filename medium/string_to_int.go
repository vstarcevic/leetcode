// String to Integer (atoi)
// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

// The algorithm for myAtoi(string s) is as follows:

// Read in and ignore any leading whitespace.
// Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
// Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
// Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
// If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
// Return the integer as the final result.

// Example 1:
// Input: s = "42"
// Output: 42
// Explanation: The underlined characters are what is read in, the caret is the current reader position.
// Step 1: "42" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "42" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "42" ("42" is read in)
//            ^
// The parsed integer is 42.
// Since 42 is in the range [-231, 231 - 1], the final result is 42.

// Example 2:
// Input: s = "   -42"
// Output: -42
// Explanation:
// Step 1: "   -42" (leading whitespace is read and ignored)
//             ^
// Step 2: "   -42" ('-' is read, so the result should be negative)
//              ^
// Step 3: "   -42" ("42" is read in)
//                ^
// The parsed integer is -42.
// Since -42 is in the range [-231, 231 - 1], the final result is -42.

package medium

import "math"

func myAtoi(s string) int {

	const (
		None   = 0
		Letter = 1
		Number = 2
		Sign   = 3
	)

	var result uint64 = 0
	isNegative := false
	lastChar := None

	for i := 0; i < len(s); i++ {

		// SPACE
		if int(s[i]) == 32 {
			if lastChar == Number || lastChar == Sign {
				break
			}
			continue
		}

		// +
		if int(s[i]) == 43 {
			if lastChar == None {
				lastChar = Sign
			} else if lastChar == Number {
				break
			} else {
				return 0
			}
			continue
		}

		// -
		if int(s[i]) == 45 {
			if lastChar == None {
				lastChar = Sign
				isNegative = true
				continue
			} else if lastChar == Number {
				break
			} else {
				return 0
			}

		}

		// .
		if int(s[i]) == 46 {
			if lastChar == None || lastChar == Number {
				break
			} else {
				return 0
			}
		}

		// 0 - 9
		if int(s[i]) >= 48 && int(s[i]) <= 57 {
			lastChar = Number
			result = result*10 + uint64(s[i]) - 48
			if result > 1<<31 {
				break
			}
			continue
		}

		// letter a-z
		if lastChar == Number {
			break
		} else if lastChar == None {
			return 0
		} else if lastChar == Sign {
			return 0
		}

	}

	if isNegative {
		return int(math.Min(float64(result), 1<<31)) * -1
	}

	return int(math.Min(float64(result), 1<<31-1))

}
