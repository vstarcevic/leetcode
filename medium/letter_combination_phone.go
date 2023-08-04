// Letter Combinations of a Phone Number
// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// 1 - nothing
// 2 - abc
// 3 - def
// 4 - ghi
// 5 - jkl
// 6 - mno
// 7 - pqrs
// 8 - tuv
// 9 - wxyz

package medium

func letterCombinations(digits string) []string {

	if digits == "" {
		return []string{}
	}

	phoneMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	result := []string{""}

	for i := 0; i < len(digits); i++ {
		resultLen := len(result)

		for j := 0; j < resultLen; j++ {
			for _, r := range phoneMap[string(digits[i])] {
				result = append(result, result[0]+r)
			}
			result = result[1:]
		}

	}

	return result

}
