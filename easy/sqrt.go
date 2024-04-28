// Sqrt(x)
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
// You must not use any built-in exponent function or operator.
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

// Example 1:
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.

package easy

import "math"

func mySqrt(x int) int {

	var tmp float64 = float64(x / 2)
	var z float64 = 1.0

	for {

		z -= (float64(z*z) - float64(x)) / float64(2*z)
		tmp = z

		if math.Abs(tmp*tmp-float64(x)) < 1 {
			break
		}

	}
	return int(tmp)

}
