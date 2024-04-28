// Combination Sum

// Given an array of distinct integers candidates and a target integer target,
//  return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

// The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency
//  of at least one of the chosen numbers is different.

// The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations
// for the given input.

// Example 1:
// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.

// Example 2:
// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]

// Example 3:
// Input: candidates = [2], target = 1
// Output: []

package medium

func combinationSum(candidate []int, target int) [][]int {

	var results [][]int = make([][]int, 0)
	var curr []int = make([]int, 0)
	var backtrack func(i int, curr []int, total int)

	backtrack = func(i int, curr []int, total int) {
		if total == target {
			var temp []int = make([]int, 0)
			temp = append(temp, curr...)
			results = append(results, temp)
			return
		}
		if i >= len(candidate) || total > target {
			return
		}

		curr = append(curr, candidate[i])
		backtrack(i, curr, total+candidate[i])
		curr = curr[:len(curr)-1]
		backtrack(i+1, curr, total)

	}

	backtrack(0, curr, 0)

	return results
}
