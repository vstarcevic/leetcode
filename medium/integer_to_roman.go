// Integer to Roman
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given an integer, convert it to a roman numeral.

// Example 1:

// Input: num = 3
// Output: "III"
// Explanation: 3 is represented as 3 ones.
// Example 2:

// Input: num = 58
// Output: "LVIII"
// Explanation: L = 50, V = 5, III = 3.
// Example 3:

// Input: num = 1994
// Output: "MCMXCIV"
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

package medium

import (
	"math"
	"strings"
)

func intToRoman(num int) string {

	var length = int(math.Log10(float64(num)) + 1)
	result := make([]string, 0)

	for i := length - 1; i >= 0; i-- {
		digit := num / int(math.Pow10(i))
		num -= digit * int(math.Pow10(i))

		switch digit {
		case 0:
		case 1:
			// digit position
			if i == 3 {
				result = append(result, "M")
			} else if i == 2 {
				result = append(result, "C")
			} else if i == 1 {
				result = append(result, "X")
			} else if i == 0 {
				result = append(result, "I")
			}
		case 2:
			// digit position
			if i == 3 {
				result = append(result, "MM")
			} else if i == 2 {
				result = append(result, "CC")
			} else if i == 1 {
				result = append(result, "XX")
			} else if i == 0 {
				result = append(result, "II")
			}
		case 3:
			// digit position
			if i == 3 {
				result = append(result, "MMM")
			} else if i == 2 {
				result = append(result, "CCC")
			} else if i == 1 {
				result = append(result, "XXX")
			} else if i == 0 {
				result = append(result, "III")
			}
		case 4:
			// digit position
			if i == 3 {
				result = append(result, "-")
			} else if i == 2 {
				result = append(result, "CD")
			} else if i == 1 {
				result = append(result, "XL")
			} else if i == 0 {
				result = append(result, "IV")
			}
		case 5:
			// digit position
			if i == 3 {
				result = append(result, "-")
			} else if i == 2 {
				result = append(result, "D")
			} else if i == 1 {
				result = append(result, "L")
			} else if i == 0 {
				result = append(result, "V")
			}
		case 6:
			// digit position
			if i == 3 {
				result = append(result, "-")
			} else if i == 2 {
				result = append(result, "DC")
			} else if i == 1 {
				result = append(result, "LX")
			} else if i == 0 {
				result = append(result, "VI")
			}
		case 7:
			// digit position
			if i == 3 {
				result = append(result, "-")
			} else if i == 2 {
				result = append(result, "DCC")
			} else if i == 1 {
				result = append(result, "LXX")
			} else if i == 0 {
				result = append(result, "VII")
			}
		case 8:
			// digit position
			if i == 3 {
				result = append(result, "-")
			} else if i == 2 {
				result = append(result, "DCCC")
			} else if i == 1 {
				result = append(result, "LXXX")
			} else if i == 0 {
				result = append(result, "VIII")
			}
		case 9:
			// digit position
			if i == 3 {
				result = append(result, "-")
			} else if i == 2 {
				result = append(result, "CM")
			} else if i == 1 {
				result = append(result, "XC")
			} else if i == 0 {
				result = append(result, "IX")
			}

		}

	}

	return strings.Join(result, "")

}
