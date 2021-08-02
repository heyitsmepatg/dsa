package ispalindrome

import "testing"

func TestIsPalindrome_Example1(t *testing.T) {
	x := 121

	isPalindrome := isPalindrome(x)

	if !isPalindrome {
		t.Errorf("Expected: true. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_Example2(t *testing.T) {
	x := -121

	isPalindrome := isPalindrome(x)

	if isPalindrome {
		t.Errorf("Expected: false. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_Example3(t *testing.T) {
	x := 10

	isPalindrome := isPalindrome(x)

	if isPalindrome {
		t.Errorf("Expected: false. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_Example4(t *testing.T) {
	x := -101

	isPalindrome := isPalindrome(x)

	if isPalindrome {
		t.Errorf("Expected: false. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_Example5(t *testing.T) {
	x := 11

	isPalindrome := isPalindrome(x)

	if !isPalindrome {
		t.Errorf("Expected: true. Actual: %v", isPalindrome)
	}
}

/*********************************
TESTS WITHOUT CONVERTING TO STRING
**********************************/
func TestIsPalindrome_NoStrings_Example1(t *testing.T) {
	x := 121

	isPalindrome := isPalindrome_noStrings(x)

	if !isPalindrome {
		t.Errorf("Expected: true. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_NoStrings_Example2(t *testing.T) {
	x := -121

	isPalindrome := isPalindrome_noStrings(x)

	if isPalindrome {
		t.Errorf("Expected: false. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_NoStrings_Example3(t *testing.T) {
	x := 10

	isPalindrome := isPalindrome_noStrings(x)

	if isPalindrome {
		t.Errorf("Expected: false. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_NoStrings_Example4(t *testing.T) {
	x := -101

	isPalindrome := isPalindrome_noStrings(x)

	if isPalindrome {
		t.Errorf("Expected: false. Actual: %v", isPalindrome)
	}
}

func TestIsPalindrome_NoStrings_Example5(t *testing.T) {
	x := 11

	isPalindrome := isPalindrome_noStrings(x)

	if !isPalindrome {
		t.Errorf("Expected: true. Actual: %v", isPalindrome)
	}
}
