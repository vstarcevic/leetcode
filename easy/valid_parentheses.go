// Valid Parentheses
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

package easy

func isValid(s string) bool {

	stack := make([]string, 0)

	for i := 0; i < len(s); i++ {

		curr := string(s[i])

		switch curr {
		case "{":
			stack = append(stack, "{")
		case "(":
			stack = append(stack, "(")
		case "[":
			stack = append(stack, "[")
		case ")":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, ")")
			}
		case "]":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, "]")
			}
		case "}":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, "}")
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}

}
