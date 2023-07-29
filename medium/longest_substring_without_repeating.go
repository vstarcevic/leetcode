// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

package medium

func longestSubstringWithoutRepeatingCharacters(s string) int {
	// Initialize the sliding window.
	start := 0
	end := 0
	charSet := make(map[byte]bool)
	maxLen := 0

	// Iterate over the string.
	for i := 0; i < len(s); i++ {
		// If the character is not in the sliding window, add it.
		if _, ok := charSet[s[i]]; !ok {
			charSet[s[i]] = true
			end++
			maxLen = max(maxLen, end-start)
		} else {
			// If the character is in the sliding window, remove the first character.
			delete(charSet, s[start])
			start++
		}
	}

	// Return the maximum length.
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
