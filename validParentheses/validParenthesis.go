package validparentheses

// Conceptually, a set of opening brackets
var OPENING = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
}

// Conceptually, a set of closing brackets
var CLOSING = map[rune]bool{
	')': true,
	']': true,
	'}': true,
}

// Explanation of Solution:
// This problem requires utilization of a Stack to model the recursive nature of the problem.
// Approach:
// 	- Iterate through the characters of the string
// 	- If bracket is opening, add to the stack
// 	- If bracket is closing and stack is not empty, compare to the element on the top of the stack
//     - If matching, continue
//     - If not matching or stack is empty, return false
// 	- If at any point, the number of elements on the stack exceeds the number of remaining elements, return false
// 	- If you iterate through all characters without error, return false
// 	- Quick catches: return if odd length of string
// Time complexity:
// 	- Requires visiting each element once, O(N)
// 	- Space complexity: Due to use of the stack, in the worst case, where a string is passed and every element is an opening basket, the worst case space complexity is O(N). The check on if the stack is larger that remaining elements restricts this to O(N/2), however, this is still a O(N) space complexity.
func isValid(s string) bool {
	strLen := len(s)
	// If given length is odd, string is invalid
	if strLen%2 != 0 || CLOSING[rune(s[0])] {
		return false
	}

	idx := -1
	var stack []rune
	for i, r := range s {
		if OPENING[r] {
			stack = append(stack, r)
			idx++
		} else if CLOSING[r] {
			// If matching, pop opening off stack.
			// Stack must be non-empty
			if len(stack) > 0 && isMatching(stack[idx], r) {
				stack = stack[:len(stack)-1]
				idx--
			} else {
				return false
			}
		}
		// If the stack becomes longer than the number
		// of characters remaining, then a valid string
		// cannot be formed
		if len(stack) > strLen-i {
			return false
		}
	}
	return true
}

func isMatching(opening rune, closing rune) bool {
	switch opening {
	case '(':
		return closing == ')'
	case '{':
		return closing == '}'
	case '[':
		return closing == ']'
	default:
		return false
	}
}
