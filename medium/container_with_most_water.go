// Container With Most Water
// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Example 1:

// 8        | |______________________|_|_________
// 7        | |                      | |       | |
// 6        | |  | |                 | |       | |
// 5        | |  | |       | |       | |       | |
// 4        | |  | |       | |  | |  | |       | |
// 3        | |  | |       | |  | |  | |       | |
// 2        | |  | |       | |  | |  | |  | |  | |
// 1   | |	| |  | |  | |  | |  | |  | |  | |  | |
// 0   | |  | |  | |  | |  | |  | |  | |  | |  | |
//      1    2    3    4    5    6    7    8    9

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
// In this case, the max area of water (blue section) the container can contain is 49.

// Example 2:
// Input: height = [1,1]
// Output: 1
package medium

import "math"

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left != right {
		// rectangle a and b sides
		a := int(math.Min(float64(height[right]), float64(height[left])))
		b := right - left

		maxArea = int(math.Max(float64(maxArea), float64(a*b)))

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}

	}

	return maxArea
}
