package longestcommonprefix

import "math"

func longestCommonPrefix(strs []string) string {
	var longestCommonPrefix string
	shortestIdx := getShortestIdx(strs)
	for i, r := range strs[shortestIdx] {
		allMatch := true
		for _, str := range strs {
			if r != rune(str[i]) {
				allMatch = false
			}
		}
		if allMatch {
			longestCommonPrefix = strs[shortestIdx][:i+1]
		} else {
			break
		}
	}
	return longestCommonPrefix
}

// helper method returns the idx of the shortest string
func getShortestIdx(strs []string) int {
	max := math.MaxInt32
	var idx int
	for i, str := range strs {
		if len(str) < max {
			max = len(str)
			idx = i
		}
	}
	return idx
}
