package longestcommonprefix

import "testing"

// Approach:
// Find shortest string in the array, then iterate through shortestString's individual characters (aka rune in Golang). For each rune in shortestString, then iterate through the initial array and compare the same index of the rune in the shortestString.
// If all runes match, replace existing longestCommonPrefix with the newest longest prefix.
// If at any point the runes do not match, then break from the first loop and return longestCommonPrefix and voila!
// Time complexity: O(n), as we must go through each element in the array once (n = length of input array). At first glance, because there is a nested for loop, it may appear as O(n^2), but because len(array) >> len(individual string) it is only O(n). A more precise analysis might show O(2*n) due to the search for the shortest string.
// Space complexity: O(1). We only require helper temp variables.
// Performance: Great :-)
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
