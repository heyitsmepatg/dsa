package longestcommonprefix

import "testing"

func TestLongestCommonPrefix_Example1(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}

	lcp := longestCommonPrefix(strs)

	if lcp != "fl" {
		t.Errorf("Expected: fl. Actual: %v", lcp)
	}
}

func TestLongestCommonPrefix_Example2(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}

	lcp := longestCommonPrefix(strs)

	if lcp != "" {
		t.Errorf("Expected: ''. Actual: %v", lcp)
	}
}

func TestLongestCommonPrefix_Example3(t *testing.T) {
	strs := []string{"BTS"}

	lcp := longestCommonPrefix(strs)

	if lcp != "BTS" {
		t.Errorf("Expected: BTS. Actual: %v", lcp)
	}
}

func TestLongestCommonPrefix_Example4(t *testing.T) {
	strs := []string{"A"}

	lcp := longestCommonPrefix(strs)

	if lcp != "A" {
		t.Errorf("Expected: A. Actual: %v", lcp)
	}
}

func TestLongestCommonPrefix_Example5(t *testing.T) {
	strs := []string{"ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF"}

	lcp := longestCommonPrefix(strs)

	if lcp != "ABCDEF" {
		t.Errorf("Expected: ABCDEF. Actual: %v", lcp)
	}
}

func TestLongestCommonPrefix_Example6(t *testing.T) {
	strs := []string{"cir", "car"}

	lcp := longestCommonPrefix(strs)

	if lcp != "c" {
		t.Errorf("Expected: c. Actual: %v", lcp)
	}
}
