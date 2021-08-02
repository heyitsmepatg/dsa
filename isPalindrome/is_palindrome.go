package ispalindrome

import "strconv"

// Approach: Recurison. A integer is a Palindrome if the
// integer is still a palindrome after removing the first
// and last digits
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return iP(strconv.Itoa(x))
}

func iP(num string) bool {
	if len(num) == 1 || len(num) == 0 {
		return true
	}
	if num[0] == num[len(num)-1] {
		return iP(num[1 : len(num)-1])
	} else {
		return false
	}
}

// Now do it without strings
func isPalindrome_noStrings(x int) bool {
	if x < 0 {
		return false
	}
	digits := getDigits(x)
	return iP_int(digits)
}

func iP_int(digits []int) bool {
	if len(digits) == 1 || len(digits) == 0 {
		return true
	}
	if (digits[0]) == digits[len(digits)-1] {
		return iP_int(digits[1 : len(digits)-1])
	} else {
		return false
	}
}

// Helper method to avoid strings
func getDigits(x int) []int {
	digits := make([]int, 0)
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	return digits
}
