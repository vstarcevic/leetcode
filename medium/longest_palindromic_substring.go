// Longest Palindromic Substring
// Given a string s, return the longest
// palindromic substring in s

// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.

// Example 2:
// Input: s = "cbbd"
// Output: "bb"

// Example 3:
// Input: s = "aba"
// Output: "aba"

package medium

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var l, r, pl, pr int

	for l <= len(s) {

		// if letters are the same, increase pointer
		for r+1 < len(s) && s[l] == s[r+1] {
			r++
		}

		// checking left and right for palindrome
		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}

		if r-l > pr-pl {
			pl, pr = l, r
		}

		l = (l+r)/2 + 1
		r = l

	}

	return s[pl : pr+1]

}
