// Longest Common Prefix
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

package easy

func longestCommonPrefix(strs []string) string {

	result := ""
	counter := 0

	if len(strs) == 0 {
		return ""
	}
out:
	for {

		if len(strs[0])-1 < counter {
			break
		}
		tmpLetter := string(strs[0][counter])

		for i := 0; i < len(strs); i++ {

			if len(strs[i]) > counter {
				var letter = string(strs[i][counter])

				if letter != tmpLetter {
					tmpLetter = ""
					break out
				}
			} else {
				tmpLetter = ""
			}
		}
		result += tmpLetter
		counter++
	}

	return result

}
