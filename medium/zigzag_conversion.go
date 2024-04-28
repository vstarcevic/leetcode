// Zigzag Conversion
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows:

// Example 1:
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"

// Example 2:
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

package medium

import "strings"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	var result []string

	for j := 0; j < numRows; j++ {

		for i := j; i < len(s); i += (numRows - 1) * 2 {
			result = append(result, string(s[i]))

			if j > 0 && j < numRows-1 && i+(numRows-1)*2-2*j < len(s) {
				result = append(result, string(s[i+(numRows-1)*2-2*j]))
			}

		}

	}

	return strings.Join(result, "")

}
